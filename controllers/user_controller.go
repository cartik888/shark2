package controllers

import (
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) GetDashboard(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user models.User
	if err := config.DB.Preload("Profile").Preload("Subscriptions.Plan").First(&user, userID).Error; err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	// Get subscription statistics
	var activeSubscriptions int64
	config.DB.Model(&models.Subscription{}).Where("user_id = ? AND status = ?", userID, "active").Count(&activeSubscriptions)

	var totalPayments float64
	config.DB.Model(&models.Payment{}).Where("user_id = ? AND status = ?", userID, "completed").Select("COALESCE(SUM(amount), 0)").Scan(&totalPayments)

	var recentPayments []models.Payment
	config.DB.Preload("Subscription.Plan").Where("user_id = ?", userID).Order("created_at DESC").Limit(5).Find(&recentPayments)

	dashboardData := gin.H{
		"user":                 user,
		"active_subscriptions": activeSubscriptions,
		"total_payments":       totalPayments,
		"recent_payments":      recentPayments,
	}

	utils.SuccessResponse(c, "Dashboard data retrieved successfully", dashboardData)
}

func (uc *UserController) GetSubscriptions(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var subscriptions []models.Subscription
	if err := config.DB.Preload("Plan").Where("user_id = ?", userID).Find(&subscriptions).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve subscriptions", err)
		return
	}

	utils.SuccessResponse(c, "Subscriptions retrieved successfully", subscriptions)
}

func (uc *UserController) GetPayments(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var payments []models.Payment
	if err := config.DB.Preload("Subscription.Plan").Where("user_id = ?", userID).Order("created_at DESC").Find(&payments).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve payments", err)
		return
	}

	utils.SuccessResponse(c, "Payments retrieved successfully", payments)
}

func (uc *UserController) BuyPlan(c *gin.Context) {
	userIDVal, _ := c.Get("user_id")
	uid := userIDVal.(uint) // extract once

	var req struct {
		PlanID          uint   `json:"plan_id" binding:"required"`
		PaymentMethodID *uint  `json:"payment_method_id"`
		CredKey         string `json:"cred_key"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	// Validate plan exists
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
	if err := config.DB.Where("user_id = ? AND plan_id = ? AND status = ?", uid, req.PlanID, "active").First(&existingSubscription).Error; err == nil {
		utils.ValidationErrorResponse(c, "You already have an active subscription for this plan")
		return
	}

	// Handle credential key if provided
	var discount float64 = 0
	if req.CredKey != "" {
		var credKey models.CredKey
		if err := config.DB.Where("key_value = ? AND is_active = ?", req.CredKey, true).First(&credKey).Error; err != nil {
			utils.NotFoundResponse(c, "Invalid credential key")
			return
		}

		if credKey.UsedBy != nil {
			utils.ValidationErrorResponse(c, "Credential key already used")
			return
		}

		// Apply discount based on key type
		switch credKey.KeyType {
		case "trial":
			discount = plan.Price * 0.5 // 50% discount
		case "partner":
			discount = plan.Price * 0.3 // 30% discount
		case "activation":
			discount = plan.Price * 0.1 // 10% discount
		}

		// Mark key as used
		credKey.UsedBy = &uid
		credKey.CurrentUses++
		config.DB.Save(&credKey)
	}

	finalAmount := plan.Price - discount

	// Create subscription with proper dates
	now := time.Now()
	var endDate *time.Time

	// Calculate end date based on plan interval
	switch plan.Interval {
	case "monthly":
		endTime := now.AddDate(0, 1, 0) // Add 1 month
		endDate = &endTime
	case "yearly":
		endTime := now.AddDate(1, 0, 0) // Add 1 year
		endDate = &endTime
	case "weekly":
		endTime := now.AddDate(0, 0, 7) // Add 7 days
		endDate = &endTime
	default:
		endTime := now.AddDate(0, 1, 0) // Default to 1 month
		endDate = &endTime
	}

	subscription := models.Subscription{
		UserID:          uid,
		PlanID:          req.PlanID,
		Status:          "active",
		PaymentMethodID: req.PaymentMethodID,
		AutoRenew:       true,
		StartDate:       &now,    // Now using pointer
		EndDate:         endDate, // Already a pointer
	}

	if err := config.DB.Create(&subscription).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create subscription", err)
		return
	}

	// Create payment record
	payment := models.Payment{
		UserID:         uid,
		SubscriptionID: &subscription.ID,
		Amount:         finalAmount,
		Currency:       plan.Currency,
		Status:         "completed",
		PaymentMethod:  "stripe",
		Description:    "Subscription payment for " + plan.Name,
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create payment record", err)
		return
	}

	// Generate subscription key
	subscriptionKey := models.SubscriptionKey{
		OriginalKeyID:    plan.ID, // or req.PlanID, depending on your logic
		DummyKey:         utils.GenerateSubscriptionKey(uid),
		AssignedToUserID: &uid,
		IsUsed:           false,
	}

	if err := config.DB.Create(&subscriptionKey).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create subscription key", err)
		return
	}

	// Load subscription with plan
	config.DB.Preload("Plan").First(&subscription, subscription.ID)

	utils.SuccessResponse(c, "Plan purchased successfully", gin.H{
		"subscription": subscription,
		"payment":      payment,
		"key":          subscriptionKey,
	})
}

func (uc *UserController) GetPlanUsage(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var subscriptions []models.Subscription
	if err := config.DB.Preload("Plan").Where("user_id = ? AND status = ?", userID, "active").Find(&subscriptions).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve subscriptions", err)
		return
	}

	var usage []gin.H
	for _, sub := range subscriptions {
		// Get subscription key (assuming OriginalKeyID is PlanID and AssignedToUserID is userID)
		var key models.SubscriptionKey
		config.DB.Where("original_key_id = ? AND assigned_to_user_id = ?", sub.PlanID, userID).First(&key)

		usage = append(usage, gin.H{
			"subscription": sub,
			"key":          key,
			"is_used":      key.IsUsed,
		})
	}

	utils.SuccessResponse(c, "Plan usage retrieved successfully", usage)
}

// Get MY subscription keys after payment
func (uc *UserController) GetMyKeys(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Debug: Check what's in the database
	var debugKeys []models.SubscriptionKey
	config.DB.Find(&debugKeys)

	// Get all subscription keys for this user
	var subscriptionKeys []models.SubscriptionKey
	if err := config.DB.Where("assigned_to_user_id = ?", userID).Find(&subscriptionKeys).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve your keys", err)
		return
	}

	// If no keys found, return debug info
	if len(subscriptionKeys) == 0 {
		utils.SuccessResponse(c, "No subscription keys found", gin.H{
			"keys":       []gin.H{},
			"total_keys": 0,
			"debug_info": gin.H{
				"user_id":          userID,
				"total_keys_in_db": len(debugKeys),
				"all_keys_in_db":   debugKeys,
			},
		})
		return
	}

	// Get keys with plan details
	var myKeys []gin.H
	for _, key := range subscriptionKeys {
		var subscription models.Subscription
		var plan models.Plan

		// Get subscription and plan details
		config.DB.Where("user_id = ? AND plan_id = ?", userID, key.OriginalKeyID).First(&subscription)
		config.DB.First(&plan, key.OriginalKeyID)

		myKeys = append(myKeys, gin.H{
			"key":         key.DummyKey,
			"key_id":      key.ID,
			"plan_name":   plan.Name,
			"plan_price":  plan.Price,
			"is_used":     key.IsUsed,
			"assigned_at": key.AssignedAt,
			"status":      subscription.Status,
		})
	}

	utils.SuccessResponse(c, "Your subscription keys retrieved successfully", gin.H{
		"keys":       myKeys,
		"total_keys": len(myKeys),
	})
}

// Get specific key by payment ID
func (uc *UserController) GetKeyByPayment(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req struct {
		PaymentID uint `json:"payment_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Payment ID is required")
		return
	}

	// Get payment details
	var payment models.Payment
	if err := config.DB.Preload("Subscription.Plan").Where("id = ? AND user_id = ?", req.PaymentID, userID).First(&payment).Error; err != nil {
		utils.NotFoundResponse(c, "Payment not found")
		return
	}

	// Get subscription key for this payment
	var subscriptionKey models.SubscriptionKey
	if err := config.DB.Where("original_key_id = ? AND assigned_to_user_id = ?", payment.Subscription.PlanID, userID).First(&subscriptionKey).Error; err != nil {
		utils.NotFoundResponse(c, "Subscription key not found for this payment")
		return
	}

	utils.SuccessResponse(c, "Your subscription key retrieved successfully", gin.H{
		"key":          subscriptionKey.DummyKey,
		"key_id":       subscriptionKey.ID,
		"payment":      payment,
		"subscription": payment.Subscription,
		"plan":         payment.Subscription.Plan,
		"is_used":      subscriptionKey.IsUsed,
		"assigned_at":  subscriptionKey.AssignedAt,
	})
}
