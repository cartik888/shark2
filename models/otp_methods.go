package models

import (
	"math/rand"
	"subscription-saas-backend/config"
	"time"
)

// GenerateOTP creates a 6-digit OTP, stores it in the database, and returns it
func GenerateOTP(email string, expiryMinutes int) string {
	rand.Seed(time.Now().UnixNano())
	otp := ""
	for i := 0; i < 6; i++ {
		otp += string('0' + rand.Intn(10))
	}

	// Mark previous OTPs as used
	config.DB.Model(&OTP{}).Where("email = ? AND used = ?", email, false).Update("used", true)

	otpEntry := OTP{
		Email:     email,
		Code:      otp,
		ExpiresAt: time.Now().Add(time.Duration(expiryMinutes) * time.Minute),
		Used:      false,
		CreatedAt: time.Now(),
	}
	config.DB.Create(&otpEntry)

	return otp
}

// VerifyOTP checks the OTP from the database for the given email, but does not mark it as used
func VerifyOTP(email, otp string) bool {
	var otpEntry OTP
	err := config.DB.Where("email = ? AND code = ? AND used = ?", email, otp, false).First(&otpEntry).Error
	if err != nil {
		return false
	}
	if otpEntry.ExpiresAt.Before(time.Now()) {
		return false
	}
	return true
}

// MarkOTPUsed marks the OTP as used after successful password reset
func MarkOTPUsed(email, otp string) {
	config.DB.Model(&OTP{}).Where("email = ? AND code = ? AND used = ?", email, otp, false).Update("used", true)
}
