package models

import (
	"fmt"
	"subscription-saas-backend/utils"
	"time"
)

type SubscriptionKeyUsage struct {
	ID                uint       `json:"id" gorm:"primaryKey"`
	SubscriptionKeyID uint       `json:"subscription_key_id" gorm:"not null"`
	UserID            uint       `json:"user_id" gorm:"not null"`
	UsedMinutes       uint       `json:"used_minutes" gorm:"default:0"`
	LastUsedAt        *time.Time `json:"last_used_at"`

	SubscriptionKey SubscriptionKey `json:"subscription_key" gorm:"foreignKey:SubscriptionKeyID;constraint:OnDelete:CASCADE"`
	User            User            `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

// Syncs this usage record to Firestore
func (sku *SubscriptionKeyUsage) SyncToFirestore() error {
	userID := sku.UserID
	planDurationMinutes := sku.UsedMinutes // or get from plan if needed
	secondsLeft := int64(planDurationMinutes * 60)
	subscriptionKey := sku.SubscriptionKey.DummyKey // use DummyKey field
	return utils.SetUserSubscription(fmt.Sprintf("%d", userID), secondsLeft, subscriptionKey)
}
