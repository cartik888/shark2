package models

import (
	"time"
)

type SubscriptionKey struct {
	ID               uint       `json:"id" gorm:"primaryKey"`
	OriginalKeyID    uint       `json:"original_key_id" gorm:"not null"`
	DummyKey         string     `json:"dummy_key" gorm:"type:varchar(191);unique;not null"`
	AssignedToUserID *uint      `json:"assigned_to_user_id"` // Pointer to allow null
	IsUsed           bool       `json:"is_used" gorm:"default:false"`
	AssignedAt       *time.Time `json:"assigned_at"` // Pointer to allow null

	// Relationships
	OriginalKey    CredKey `json:"original_key" gorm:"foreignKey:OriginalKeyID;constraint:OnDelete:CASCADE"`
	AssignedToUser *User   `json:"assigned_to_user,omitempty" gorm:"foreignKey:AssignedToUserID;constraint:OnDelete:SET NULL"`
}
