package utils

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

var client *firestore.Client

// Call this once in your main function
func InitFirestore(projectID string) error {
	var err error
	client, err = firestore.NewClient(context.Background(), projectID)
	return err
}

// Set user subscription time and key in Firestore
func SetUserSubscription(userID string, secondsLeft int64, subscriptionKey string) error {
	if client == nil {
		return ErrFirestoreNotInitialized
	}
	_, err := client.Collection("user_times").Doc(userID).Set(context.Background(), map[string]interface{}{
		"seconds_left":     secondsLeft,
		"subscription_key": subscriptionKey,
	}, firestore.MergeAll)
	if err != nil {
		log.Printf("[ERROR] Firestore write failed: %v", err)
		return err
	}
	log.Printf("[DEBUG] Firestore write success for user %s", userID)
	return nil
}

func GetUserTimeFromFirestore(ctx context.Context, userID uint) (int, error) {
	if client == nil {
		return 0, ErrFirestoreNotInitialized
	}
	doc, err := client.Collection("user_times").Doc(fmt.Sprintf("%d", userID)).Get(ctx)
	if err != nil {
		return 0, err
	}
	val, ok := doc.Data()["seconds_left"].(int64)
	if !ok {
		return 0, nil
	}
	return int(val), nil
}

var ErrFirestoreNotInitialized = &FirestoreError{"Firestore client not initialized"}

type FirestoreError struct {
	msg string
}

func (e *FirestoreError) Error() string {
	return e.msg
}
