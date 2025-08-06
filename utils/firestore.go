package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

var client *firestore.Client

// Call this once in your main function
func InitFirestore(projectID string) error {
	// Step 1: Get base64 encoded credentials from environment variable
	credsB64 := os.Getenv("GOOGLE_CREDENTIALS_B64")
	if credsB64 == "" {
		return fmt.Errorf("GOOGLE_CREDENTIALS_B64 env var not set")
	}

	// Step 2: Decode base64 string
	decoded, err := base64.StdEncoding.DecodeString(credsB64)
	if err != nil {
		return fmt.Errorf("failed to decode GOOGLE_CREDENTIALS_B64: %v", err)
	}

	// Step 3: Write decoded content to a temporary JSON file
	tmpFile, err := os.CreateTemp("", "firebase-creds-*.json")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %v", err)
	}
	defer tmpFile.Close()

	if _, err := tmpFile.Write(decoded); err != nil {
		return fmt.Errorf("failed to write decoded credentials: %v", err)
	}

	// Step 4: Set GOOGLE_APPLICATION_CREDENTIALS to point to the temp file
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", tmpFile.Name())

	// Step 5: Initialize Firestore client
	ctx := context.Background()
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("failed to initialize Firestore: %v", err)
	}

	log.Println("✅ Firestore initialized successfully")
	return nil
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

// Get user subscription time from Firestore
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

// Error type for uninitialized Firestore client
var ErrFirestoreNotInitialized = &FirestoreError{"Firestore client not initialized"}

type FirestoreError struct {
	msg string
}

func (e *FirestoreError) Error() string {
	return e.msg
}
