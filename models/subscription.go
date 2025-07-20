package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserID          uint           `json:"user_id" gorm:"not null;index"`
	PlanID          uint           `json:"plan_id" gorm:"not null;index"`
	Status          string         `json:"status" gorm:"type:enum('active','inactive','cancelled','expired');default:'active'"`
	PaymentMethodID *uint          `json:"payment_method_id" gorm:"index"`
	AutoRenew       bool           `json:"auto_renew" gorm:"default:true"`
	StartDate       *time.Time     `json:"start_date"`
	EndDate         *time.Time     `json:"end_date"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	User          User          `json:"user" gorm:"foreignKey:UserID"`
	Plan          Plan          `json:"plan" gorm:"foreignKey:PlanID"`
	PaymentMethod PaymentMethod `json:"payment_method" gorm:"foreignKey:PaymentMethodID"`
	Payments      []Payment     `json:"payments" gorm:"foreignKey:SubscriptionID"`
}
