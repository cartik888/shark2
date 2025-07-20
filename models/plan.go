package models

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Price       float64        `json:"price" gorm:"not null"`
	Currency    string         `json:"currency" gorm:"default:'USD'"`
	Interval    string         `json:"interval" gorm:"default:'monthly'"`
	Features    string         `json:"features"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Subscriptions []Subscription `json:"subscriptions,omitempty" gorm:"foreignKey:PlanID"`
}
