package controllers

import (
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"

	"github.com/gin-gonic/gin"
)

type PaymentController struct{}

func (pc *PaymentController) GetPayments(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var payments []models.Payment
	if err := config.DB.Preload("Subscription.Plan").Where("user_id = ?", userID).Order("created_at DESC").Find(&payments).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve payments", err)
		return
	}

	utils.SuccessResponse(c, "Payments retrieved successfully", payments)
}

func (pc *PaymentController) GetPayment(c *gin.Context) {
	userID, _ := c.Get("user_id")
	paymentID := c.Param("id")

	var payment models.Payment
	if err := config.DB.Preload("Subscription.Plan").Where("id = ? AND user_id = ?", paymentID, userID).First(&payment).Error; err != nil {
		utils.NotFoundResponse(c, "Payment not found")
		return
	}

	utils.SuccessResponse(c, "Payment retrieved successfully", payment)
}

func (pc *PaymentController) CreatePaymentMethod(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req struct {
		Type           string `json:"type" binding:"required"`
		Provider       string `json:"provider" binding:"required"`
		ProviderID     string `json:"provider_id" binding:"required"`
		Last4          string `json:"last4"`
		Brand          string `json:"brand"`
		ExpiryMonth    int    `json:"expiry_month"`
		ExpiryYear     int    `json:"expiry_year"`
		BillingAddress string `json:"billing_address"`
		IsDefault      bool   `json:"is_default"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	// If this is set as default, unset other default payment methods
	if req.IsDefault {
		config.DB.Model(&models.PaymentMethod{}).Where("user_id = ?", userID).Update("is_default", false)
	}

	paymentMethod := models.PaymentMethod{
		UserID:         userID.(uint),
		Type:           req.Type,
		Provider:       req.Provider,
		ProviderID:     req.ProviderID,
		Last4:          req.Last4,
		Brand:          req.Brand,
		ExpiryMonth:    req.ExpiryMonth,
		ExpiryYear:     req.ExpiryYear,
		BillingAddress: req.BillingAddress,
		IsDefault:      req.IsDefault,
		IsActive:       true,
	}

	if err := config.DB.Create(&paymentMethod).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create payment method", err)
		return
	}

	utils.SuccessResponse(c, "Payment method created successfully", paymentMethod)
}

func (pc *PaymentController) GetPaymentMethods(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var paymentMethods []models.PaymentMethod
	if err := config.DB.Where("user_id = ? AND is_active = ?", userID, true).Find(&paymentMethods).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve payment methods", err)
		return
	}

	utils.SuccessResponse(c, "Payment methods retrieved successfully", paymentMethods)
}
