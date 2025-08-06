package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateCredKey(keyType string) string {
	timestamp := time.Now().Unix()
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)
	randomHex := hex.EncodeToString(randomBytes)

	prefix := getKeyPrefix(keyType)
	return fmt.Sprintf("%s-%d-%s", prefix, timestamp, randomHex)
}

func GenerateSubscriptionKey(userID uint) string {
	timestamp := time.Now().Unix()
	randomBytes := make([]byte, 12)
	rand.Read(randomBytes)
	randomHex := hex.EncodeToString(randomBytes)

	return fmt.Sprintf("SUB-%d-%d-%s", userID, timestamp, randomHex)
}

// GenerateUniqueSubscriptionKey generates a secure, unique key for a subscription/payment
func GenerateUniqueSubscriptionKey(paymentID uint, userID uint) string {
	timestamp := time.Now().Unix()
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	randomHex := hex.EncodeToString(randomBytes)
	return fmt.Sprintf("SUBPAY-%d-%d-%d-%s", paymentID, userID, timestamp, randomHex)
}

func getKeyPrefix(keyType string) string {
	switch keyType {
	case "trial":
		return "TRIAL"
	case "partner":
		return "PARTNER"
	case "activation":
		return "ACTIVATE"
	default:
		return "KEY"
	}
}

func GenerateInvoiceNumber() string {
	now := time.Now()
	return fmt.Sprintf("INV-%s-%02d%02d",
		now.Format("20060102"),
		now.Hour(),
		now.Minute())
}
