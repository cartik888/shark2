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

// UPDATED CHECKOUT ENDPOINT FOR YOUR PAYLOAD
func (pc *PaymentController) CreateCheckout(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req struct {
		PlanID        uint   `json:"plan_id" binding:"required"`
		PaymentMethod string `json:"payment_method" binding:"required"`
		CredKey       string `json:"cred_key"`
		Currency      string `json:"currency"`
		SuccessURL    string `json:"success_url"`
		CancelURL     string `json:"cancel_url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	// Validate plan exists and is active
	var plan models.Plan
	if err := config.DB.First(&plan, req.PlanID).Error; err != nil {
		utils.NotFoundResponse(c, "Plan not found")
		return
	}

	if !plan.IsActive {
		utils.ValidationErrorResponse(c, "Plan is not active")
		return
	}

	// Check if user already has an active subscription for this plan
	var existingSubscription models.Subscription
	if err := config.DB.Where("user_id = ? AND plan_id = ? AND status = ?", userID, req.PlanID, "active").First(&existingSubscription).Error; err == nil {
		utils.ValidationErrorResponse(c, "You already have an active subscription for this plan")
		return
	}

	// Get user details
	var user models.User
	if err := config.DB.Preload("Profile").First(&user, userID).Error; err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	// Calculate pricing with discounts
	finalAmount := plan.Price
	discount := 0.0
	discountType := ""

	// Handle credential key if provided
	if req.CredKey != "" {
		var credKey models.CredKey
		if err := config.DB.Where("key_value = ? AND is_active = ?", req.CredKey, true).First(&credKey).Error; err != nil {
			utils.NotFoundResponse(c, "Invalid credential key")
			return
		}

		// Check if key is already used
		if credKey.MaxUses > 0 && credKey.CurrentUses >= credKey.MaxUses {
			utils.ValidationErrorResponse(c, "Credential key has reached maximum usage limit")
			return
		}

		// Apply discount based on key type
		switch credKey.KeyType {
		case "trial":
			discount = plan.Price * 0.5 // 50% discount
			discountType = "Trial Discount (50%)"
		case "partner":
			discount = plan.Price * 0.3 // 30% discount
			discountType = "Partner Discount (30%)"
		case "activation":
			discount = plan.Price * 0.1 // 10% discount
			discountType = "Activation Discount (10%)"
		default:
			discount = plan.Price * 0.05 // 5% default discount
			discountType = "Special Discount (5%)"
		}

		finalAmount = plan.Price - discount
	}

	// Set currency
	currency := req.Currency
	if currency == "" {
		currency = plan.Currency
	}

	// Generate checkout ID and invoice number
	checkoutID := utils.GenerateSubscriptionKey(userID.(uint))
	invoiceNumber := utils.GenerateInvoiceNumber()

	// Validate payment method
	validPaymentMethods := []string{"Credit Card", "Debit Card", "Bank Transfer", "PayPal", "Stripe", "Razorpay"}
	isValidPaymentMethod := false
	for _, method := range validPaymentMethods {
		if req.PaymentMethod == method {
			isValidPaymentMethod = true
			break
		}
	}

	if !isValidPaymentMethod {
		utils.ValidationErrorResponse(c, "Invalid payment method")
		return
	}

	// Create checkout session data
	checkoutData := gin.H{
		"checkout_id":    checkoutID,
		"plan":           plan,
		"user":           user,
		"original_price": plan.Price,
		"discount":       discount,
		"discount_type":  discountType,
		"final_amount":   finalAmount,
		"currency":       currency,
		"payment_method": req.PaymentMethod,
		"invoice_number": invoiceNumber,
		"success_url":    req.SuccessURL,
		"cancel_url":     req.CancelURL,
		"status":         "pending",
		"created_at":     utils.GetCurrentTime(),
		"metadata": gin.H{
			"user_id":        userID,
			"plan_id":        req.PlanID,
			"cred_key":       req.CredKey,
			"payment_method": req.PaymentMethod,
		},
	}

	// For Bank Transfer, provide additional instructions
	if req.PaymentMethod == "Bank Transfer" {
		checkoutData["bank_details"] = gin.H{
			"account_name":   "Shark SaaS Ltd",
			"account_number": "1234567890",
			"routing_number": "123456789",
			"bank_name":      "Example Bank",
			"swift_code":     "EXAMPLEXXX",
			"reference":      checkoutID,
			"instructions":   "Please use the checkout ID as reference when making the transfer",
		}
	}

	utils.SuccessResponse(c, "Checkout session created successfully", checkoutData)
}

// Process payment after successful checkout
func (pc *PaymentController) ProcessCheckout(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req struct {
		CheckoutID      string  `json:"checkout_id" binding:"required"`
		PaymentMethod   string  `json:"payment_method" binding:"required"`
		TransactionID   string  `json:"transaction_id"`
		StripePaymentID string  `json:"stripe_payment_id"`
		Status          string  `json:"status" binding:"required"`
		PlanID          uint    `json:"plan_id" binding:"required"`
		Amount          float64 `json:"amount" binding:"required"`
		Currency        string  `json:"currency"`
		CredKey         string  `json:"cred_key"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	// Validate plan
	var plan models.Plan
	if err := config.DB.First(&plan, req.PlanID).Error; err != nil {
		utils.NotFoundResponse(c, "Plan not found")
		return
	}

	// Handle credential key if provided
	if req.CredKey != "" {
		var credKey models.CredKey
		if err := config.DB.Where("key_value = ? AND is_active = ?", req.CredKey, true).First(&credKey).Error; err == nil {
			// Mark key as used
			credKey.CurrentUses++
			if credKey.MaxUses == 1 {
				userIDUint := userID.(uint)
				credKey.UsedBy = &userIDUint
			}
			config.DB.Save(&credKey)
		}
	}

	// Create subscription
	subscription := models.Subscription{
		UserID:    userID.(uint),
		PlanID:    req.PlanID,
		Status:    "active",
		AutoRenew: true,
	}

	if err := config.DB.Create(&subscription).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create subscription", err)
		return
	}

	// Create payment record
	payment := models.Payment{
		UserID:          userID.(uint),
		SubscriptionID:  &subscription.ID,
		Amount:          req.Amount,
		Currency:        req.Currency,
		Status:          req.Status,
		PaymentMethod:   req.PaymentMethod,
		TransactionID:   req.TransactionID,
		StripePaymentID: req.StripePaymentID,
		Description:     "Subscription payment for " + plan.Name,
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create payment record", err)
		return
	}

	// Generate subscription key
	subscriptionKey := models.SubscriptionKey{
		OriginalKeyID:    plan.ID,
		DummyKey:         utils.GenerateSubscriptionKey(userID.(uint)),
		AssignedToUserID: &[]uint{userID.(uint)}[0],
		IsUsed:           false,
	}

	if err := config.DB.Create(&subscriptionKey).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create subscription key", err)
		return
	}

	// Load subscription with plan
	config.DB.Preload("Plan").First(&subscription, subscription.ID)

	utils.SuccessResponse(c, "Payment processed successfully", gin.H{
		"subscription": subscription,
		"payment":      payment,
		"key":          subscriptionKey,
	})
}
