package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID                   uint           `json:"id" gorm:"primaryKey"`
	UserID               uint           `json:"user_id" gorm:"not null"`
	PlanID               uint           `json:"plan_id" gorm:"not null"`
	Status               string         `json:"status" gorm:"default:'active'"` // active, cancelled, expired, suspended
	StartDate            time.Time      `json:"start_date"`
	EndDate              *time.Time     `json:"end_date"`
	NextBillingDate      *time.Time     `json:"next_billing_date"`
	CancelledAt          *time.Time     `json:"cancelled_at"`
	TrialEndsAt          *time.Time     `json:"trial_ends_at"`
	IsTrialActive        bool           `json:"is_trial_active" gorm:"default:false"`
	AutoRenew            bool           `json:"auto_renew" gorm:"default:true"`
	PaymentMethodID      *uint          `json:"payment_method_id"`
	StripeCustomerID     string         `json:"stripe_customer_id"`
	StripeSubscriptionID string         `json:"stripe_subscription_id"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User          User                `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Plan          Plan                `json:"plan,omitempty" gorm:"foreignKey:PlanID"`
	PaymentMethod *PaymentMethod      `json:"payment_method,omitempty" gorm:"foreignKey:PaymentMethodID"`
	Events        []SubscriptionEvent `json:"events,omitempty" gorm:"foreignKey:SubscriptionID"`
	Payments      []Payment           `json:"payments,omitempty" gorm:"foreignKey:SubscriptionID"`
}

type SubscriptionEvent struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	SubscriptionID uint      `json:"subscription_id" gorm:"not null"`
	EventType      string    `json:"event_type" gorm:"not null"` // created, updated, cancelled, renewed, payment_failed
	Description    string    `json:"description"`
	Metadata       string    `json:"metadata"`
	CreatedAt      time.Time `json:"created_at"`

	// Relationships
	Subscription Subscription `json:"-" gorm:"foreignKey:SubscriptionID"`
}
