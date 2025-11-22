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

	type UserTime struct {
		SecondsLeft     int64  `datastore:"seconds_left"`
		SubscriptionKey string `datastore:"subscription_key"`
	}

	key := datastore.NameKey("user_times", userID, nil)
	_, err := client.Put(context.Background(), key, &UserTime{
		SecondsLeft:     secondsLeft,
		SubscriptionKey: subscriptionKey,
	})

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

	var data struct {
		SecondsLeft int64 `datastore:"seconds_left"`
	}

	err := client.Get(ctx, key, &data)
	if err != nil {
		return 0, err
	}

	return int(data.SecondsLeft), nil
}

var ErrDatastoreNotInitialized = &DatastoreError{"Datastore client not initialized"}

type DatastoreError struct {
	msg string
}

func (e *DatastoreError) Error() string {
	return e.msg
}
