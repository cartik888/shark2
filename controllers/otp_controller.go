package controllers

import (
    "net/http"
    "subscription-saas-backend/utils"
    "subscription-saas-backend/models"
    "github.com/gin-gonic/gin"
)

type OTPController struct{}

type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required"`
}

// POST /api/v1/send-otp
func (o *OTPController) SendOTP(c *gin.Context) {
	var req SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	otp := models.GenerateOTP(req.Email, 5)

	// 🚀 Send OTP in background (non-blocking)
	go func() {
		err := utils.SendEmailOTP(req.Email, otp)
		if err != nil {
			// log only, do not block API
			println("[OTP ERROR] Unable to send email:", err.Error())
		}
	}()

	utils.SuccessResponse(c, "OTP sent successfully", nil)
}

// POST /api/v1/verify-otp
func (o *OTPController) VerifyOTP(c *gin.Context) {
	var req VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	if models.VerifyOTP(req.Email, req.OTP) {
		utils.SuccessResponse(c, "OTP verified successfully", nil)
	} else {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired OTP", nil)
	}
}
