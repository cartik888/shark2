package models

import (
	"time"

	"gorm.io/gorm"
)

type CredKey struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	KeyValue    string         `json:"key_value" gorm:"type:varchar(191);uniqueIndex;not null"`
	KeyType     string         `json:"key_type" gorm:"not null"` // activation, trial, partner, etc.
	Description string         `json:"description"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	UsedBy      *uint          `json:"used_by"` // User ID who used this key
	UsedAt      *time.Time     `json:"used_at"`
	ExpiresAt   *time.Time     `json:"expires_at"`
	MaxUses     int            `json:"max_uses" gorm:"default:1"`
	CurrentUses int            `json:"current_uses" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User *User `json:"user,omitempty" gorm:"foreignKey:UsedBy"`
}
