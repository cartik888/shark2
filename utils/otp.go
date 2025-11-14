package utils

import (
	"math/rand"
	"sync"
	"time"
)

type otpEntry struct {
	OTP     string
	Expires time.Time
}

var (
	otpStore = make(map[string]otpEntry)
	otpMutex sync.Mutex
)

func GenerateOTP(email string, expiryMinutes int) string {
	rand.Seed(time.Now().UnixNano())

	otp := ""
	for i := 0; i < 6; i++ {
		otp += string('0' + rand.Intn(10))
	}

	otpMutex.Lock()
	otpStore[email] = otpEntry{
		OTP:     otp,
		Expires: time.Now().Add(time.Duration(expiryMinutes) * time.Minute),
	}
	otpMutex.Unlock()

	return otp
}

func VerifyOTP(email, otp string) bool {
	otpMutex.Lock()
	entry, exists := otpStore[email]
	otpMutex.Unlock()

	if !exists {
		return false
	}
	if entry.Expires.Before(time.Now()) {
		return false
	}
	if entry.OTP != otp {
		return false
	}

	otpMutex.Lock()
	delete(otpStore, email)
	otpMutex.Unlock()

	return true
}
