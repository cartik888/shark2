package controllers

import (
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"

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
	config.DB.Where("user_id = ?", userID).Order("created_at DESC").Limit(5).Find(&recentPayments)

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

	// Create subscription
	subscription := models.Subscription{
		UserID:          uid,
		PlanID:          req.PlanID,
		Status:          "active",
		PaymentMethodID: req.PaymentMethodID,
		AutoRenew:       true,
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
		config.DB.Where("Original_key_id = ? AND assigned_to_user_id = ?", sub.PlanID, userID).First(&key)

		usage = append(usage, gin.H{
			"subscription": sub,
			"key":          key,
			"is_used":      key.IsUsed,
		})
	}

	utils.SuccessResponse(c, "Plan usage retrieved successfully", usage)
}
