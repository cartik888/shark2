package models

import (
	"time"

	"gorm.io/gorm"
)

type SupportTicket struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	Subject     string         `json:"subject" gorm:"not null"`
	Description string         `json:"description" gorm:"type:text"`
	Status      string         `json:"status" gorm:"default:'open'"`     // open, in_progress, resolved, closed
	Priority    string         `json:"priority" gorm:"default:'medium'"` // low, medium, high, urgent
	Category    string         `json:"category"`                         // billing, technical, general
	AssignedTo  *uint          `json:"assigned_to"`                      // Admin user ID
	ResolvedAt  *time.Time     `json:"resolved_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User     User             `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Admin    *User            `json:"admin,omitempty" gorm:"foreignKey:AssignedTo"`
	Messages []SupportMessage `json:"messages,omitempty" gorm:"foreignKey:TicketID"`
}

type SupportMessage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	TicketID  uint      `json:"ticket_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Message   string    `json:"message" gorm:"type:text;not null"`
	IsAdmin   bool      `json:"is_admin" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Ticket SupportTicket `json:"-" gorm:"foreignKey:TicketID"`
	User   User          `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
