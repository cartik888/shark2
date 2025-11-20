package config

import "os"

var RazorpayKeyID string
var RazorpayKeySecret string

func LoadRazorpayConfig() {
	RazorpayKeyID = os.Getenv("RAZORPAY_KEY_ID")
	RazorpayKeySecret = os.Getenv("RAZORPAY_KEY_SECRET")
}
