package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"type:varchar(191);uniqueIndex;not null"`
	Password  string         `json:"-" gorm:"not null"`
	Role      string         `json:"role" gorm:"default:'user'"`
	IsAdmin   bool           `json:"is_admin" gorm:"default:false"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	IsBlocked bool           `json:"is_blocked" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Profile       *UserProfile    `json:"profile,omitempty" gorm:"foreignKey:UserID"`
	Subscriptions []Subscription  `json:"subscriptions,omitempty" gorm:"foreignKey:UserID"`
	Payments      []Payment       `json:"payments,omitempty" gorm:"foreignKey:UserID"`
	PasswordReset []PasswordReset `json:"-" gorm:"foreignKey:UserID"`
}

type UserProfile struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Company   string    `json:"company"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	ZipCode   string    `json:"zip_code"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	User *User `json:"-" gorm:"foreignKey:UserID"`
}

type PasswordReset struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Token     string    `json:"token" gorm:"type:varchar(191);uniqueIndex;not null"`
	ExpiresAt time.Time `json:"expires_at"`
	Used      bool      `json:"used" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`

	// Relationships
	User *User `json:"-" gorm:"foreignKey:UserID"`
}
