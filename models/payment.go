package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserID          uint           `json:"user_id" gorm:"not null"`
	SubscriptionID  *uint          `json:"subscription_id"`
	Amount          float64        `json:"amount" gorm:"not null"`
	Currency        string         `json:"currency" gorm:"default:'USD'"`
	Status          string         `json:"status" gorm:"default:'pending'"` // pending, completed, failed, refunded
	PaymentMethod   string         `json:"payment_method"`                  // stripe, paypal, bank_transfer
	TransactionID   string         `json:"transaction_id"`
	StripePaymentID string         `json:"stripe_payment_id"`
	Description     string         `json:"description"`
	FailureReason   string         `json:"failure_reason"`
	RefundedAmount  float64        `json:"refunded_amount" gorm:"default:0"`
	RefundedAt      *time.Time     `json:"refunded_at"`
	ProcessedAt     *time.Time     `json:"processed_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User         User          `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Subscription *Subscription `json:"subscription,omitempty" gorm:"foreignKey:SubscriptionID"`
	Invoice      *Invoice      `json:"invoice,omitempty" gorm:"foreignKey:PaymentID"`
}

type PaymentMethod struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	UserID         uint           `json:"user_id" gorm:"not null"`
	Type           string         `json:"type" gorm:"not null"` // card, bank_account, paypal
	Provider       string         `json:"provider"`             // stripe, paypal
	ProviderID     string         `json:"provider_id"`          // External provider ID
	Last4          string         `json:"last4"`                // Last 4 digits for cards
	Brand          string         `json:"brand"`                // visa, mastercard, etc.
	ExpiryMonth    int            `json:"expiry_month"`
	ExpiryYear     int            `json:"expiry_year"`
	IsDefault      bool           `json:"is_default" gorm:"default:false"`
	IsActive       bool           `json:"is_active" gorm:"default:true"`
	BillingAddress string         `json:"billing_address"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User          User           `json:"-" gorm:"foreignKey:UserID"`
	Subscriptions []Subscription `json:"subscriptions,omitempty" gorm:"foreignKey:PaymentMethodID"`
}
