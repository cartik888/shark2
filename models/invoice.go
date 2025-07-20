package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	UserID         uint           `json:"user_id" gorm:"not null"`
	SubscriptionID *uint          `json:"subscription_id"`
	PaymentID      *uint          `json:"payment_id"`
	InvoiceNumber  string         `json:"invoice_number" gorm:"type:varchar(191);uniqueIndex;not null"`
	Status         string         `json:"status" gorm:"default:'draft'"` // draft, sent, paid, overdue, cancelled
	Subtotal       float64        `json:"subtotal" gorm:"not null"`
	TaxAmount      float64        `json:"tax_amount" gorm:"default:0"`
	DiscountAmount float64        `json:"discount_amount" gorm:"default:0"`
	Total          float64        `json:"total" gorm:"not null"`
	Currency       string         `json:"currency" gorm:"default:'USD'"`
	DueDate        time.Time      `json:"due_date"`
	PaidAt         *time.Time     `json:"paid_at"`
	SentAt         *time.Time     `json:"sent_at"`
	Notes          string         `json:"notes"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User         User          `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Subscription *Subscription `json:"subscription,omitempty" gorm:"foreignKey:SubscriptionID"`
	Payment      *Payment      `json:"payment,omitempty" gorm:"foreignKey:PaymentID"`
}
