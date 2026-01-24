package utils

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
	"golang.org/x/oauth2/google"
)

var client *datastore.Client

// Call this once in your main function
func InitDatastore(projectID string) error {
	// Log the active service account and project ID
	creds, errCreds := google.FindDefaultCredentials(context.Background())
	if errCreds != nil {
		log.Println("Error finding default credentials:", errCreds)
	} else {
		log.Println("ACTIVE SA =", creds.ProjectID, creds)
	}
	var err error
	client, err = datastore.NewClient(context.Background(), projectID)
	return err
}

// Simple wrapper without context
func GetUserTimeFromFirestore(userID uint) (int, error) {
	return GetUserTimeFromDatastore(context.Background(), userID)
}

// Version with context
func GetUserTimeFromFirestoreWithContext(ctx context.Context, userID uint) (int, error) {
	return GetUserTimeFromDatastore(ctx, userID)
}

// Set user subscription time and key in Datastore
func SetUserSubscription(userID string, secondsLeft int64, subscriptionKey string) error {
	if client == nil {
		return ErrDatastoreNotInitialized
	}

	type Subscription struct {
		SecondsLeft     int64  `datastore:"seconds_left"`
		SubscriptionKey string `datastore:"subscription_key"`
	}
	type UserSubscriptions struct {
		SubscriptionKeys []Subscription `datastore:"subscription_keys"`
	}

	key := datastore.NameKey("user_times", userID, nil)
	var userSubs UserSubscriptions
	err := client.Get(context.Background(), key, &userSubs)
	if err != nil {
		// Handle legacy record with single subscription_key field
		if err.Error() == "datastore: cannot load field \"subscription_key\" into a \"utils.UserSubscriptions\": no such struct field" {
			// Try to load legacy struct
			type LegacyUserTime struct {
				SecondsLeft     int64  `datastore:"seconds_left"`
				SubscriptionKey string `datastore:"subscription_key"`
			}
			var legacy LegacyUserTime
			errLegacy := client.Get(context.Background(), key, &legacy)
			if errLegacy == nil {
				// Migrate legacy data to new array format
				userSubs.SubscriptionKeys = append(userSubs.SubscriptionKeys, Subscription{
					SecondsLeft:     legacy.SecondsLeft,
					SubscriptionKey: legacy.SubscriptionKey,
				})
			} else {
				log.Printf("[ERROR] Datastore legacy read failed: %v", errLegacy)
				return errLegacy
			}
		} else if err != datastore.ErrNoSuchEntity {
			log.Printf("[ERROR] Datastore read failed: %v", err)
			return err
		}
	}

	// Append new subscription
	newSub := Subscription{
		SecondsLeft:     secondsLeft,
		SubscriptionKey: subscriptionKey,
	}
	userSubs.SubscriptionKeys = append(userSubs.SubscriptionKeys, newSub)

	_, err = client.Put(context.Background(), key, &userSubs)
	if err != nil {
		log.Printf("[ERROR] Datastore write failed: %v", err)
		return err
	}

	log.Printf("[DEBUG] Datastore write success for user %s", userID)
	return nil
}

func GetUserTimeFromDatastore(ctx context.Context, userID uint) (int, error) {
	if client == nil {
		return 0, ErrDatastoreNotInitialized
	}

	key := datastore.NameKey("user_times", fmt.Sprint(userID), nil)

	// Try new schema first
	type Subscription struct {
		SecondsLeft     int64  `datastore:"seconds_left"`
		SubscriptionKey string `datastore:"subscription_key"`
	}
	type UserSubscriptions struct {
		SubscriptionKeys []Subscription `datastore:"subscription_keys"`
	}
	var userSubs UserSubscriptions
	err := client.Get(ctx, key, &userSubs)
	if err == nil && len(userSubs.SubscriptionKeys) > 0 {
		// Return the latest subscription's seconds_left
		latest := userSubs.SubscriptionKeys[len(userSubs.SubscriptionKeys)-1]
		return int(latest.SecondsLeft), nil
	}

	// Fallback to legacy schema
	var legacy struct {
		SecondsLeft int64 `datastore:"seconds_left"`
	}
	errLegacy := client.Get(ctx, key, &legacy)
	if errLegacy != nil {
		return 0, errLegacy
	}
	return int(legacy.SecondsLeft), nil
}

var ErrDatastoreNotInitialized = &DatastoreError{"Datastore client not initialized"}

type DatastoreError struct {
	msg string
}

func (e *DatastoreError) Error() string {
	return e.msg
}
