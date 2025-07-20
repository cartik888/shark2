package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

func GenerateActivationKey() string {
	timestamp := time.Now().Format("2006")
	randomPart := generateRandomString(8)
	return fmt.Sprintf("SHARK-ACTIVATION-%s-%s", timestamp, randomPart)
}

func GenerateSubscriptionKey(userID uint) string {
	// Generate random bytes
	bytes := make([]byte, 8)
	rand.Read(bytes)

	// Create key with user ID and random component
	return fmt.Sprintf("SUB-%d-%X", userID, bytes)
}

func GenerateTrialKey() string {
	timestamp := time.Now().Format("2006")
	randomPart := generateRandomString(6)
	return fmt.Sprintf("SHARK-TRIAL-%s-%s", timestamp, randomPart)
}

func GenerateCredKey(keyType string) string {
	// Generate random bytes
	bytes := make([]byte, 4)
	rand.Read(bytes)

	// Create timestamp suffix
	timestamp := time.Now().Format("2006")
	
	// Generate sequential number (simplified)
	seqNum := time.Now().Unix() % 1000

	// Format: SHARK-TYPE-YEAR-XXX
	return fmt.Sprintf("SHARK-%s-%s-%03d", 
		strings.ToUpper(keyType), 
		timestamp, 
		seqNum)
}

func GenerateInvoiceNumber() string {
	// Format: INV-YYYYMMDD-XXXX
	now := time.Now()
	dateStr := now.Format("20060102")
	seqNum := now.Unix() % 10000

	return fmt.Sprintf("INV-%s-%04d", dateStr, seqNum)
}

func generateRandomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	rand.Read(b)
	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b)
}
