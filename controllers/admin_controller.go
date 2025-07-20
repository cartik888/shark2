package controllers

import (
	"strconv"
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (ac *AdminController) GetDashboard(c *gin.Context) {
	var stats gin.H

	// User statistics
	var totalUsers, activeUsers, blockedUsers int64
	config.DB.Model(&models.User{}).Count(&totalUsers)
	config.DB.Model(&models.User{}).Where("is_active = ?", true).Count(&activeUsers)
	config.DB.Model(&models.User{}).Where("is_blocked = ?", true).Count(&blockedUsers)

	// Subscription statistics
	var totalSubscriptions, activeSubscriptions int64
	config.DB.Model(&models.SubscriptionKey{}).Count(&totalSubscriptions)
	config.DB.Model(&models.SubscriptionKey{}).Where("status = ?", "active").Count(&activeSubscriptions)

	// Payment statistics
	var totalRevenue float64
	var totalPayments int64
	config.DB.Model(&models.Payment{}).Where("status = ?", "completed").Count(&totalPayments)
	config.DB.Model(&models.Payment{}).Where("status = ?", "completed").Select("COALESCE(SUM(amount), 0)").Scan(&totalRevenue)

	// Recent activities
	var recentUsers []models.User
	config.DB.Preload("Profile").Order("created_at DESC").Limit(5).Find(&recentUsers)

	var recentPayments []models.Payment
	config.DB.Preload("User.Profile").Order("created_at DESC").Limit(5).Find(&recentPayments)

	stats = gin.H{
		"users": gin.H{
			"total":   totalUsers,
			"active":  activeUsers,
			"blocked": blockedUsers,
		},
		"subscriptions": gin.H{
			"total":  totalSubscriptions,
			"active": activeSubscriptions,
		},
		"payments": gin.H{
			"total":   totalPayments,
			"revenue": totalRevenue,
		},
		"recent_users":    recentUsers,
		"recent_payments": recentPayments,
	}

	utils.SuccessResponse(c, "Dashboard statistics retrieved successfully", stats)
}

func (ac *AdminController) GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	offset := (page - 1) * limit

	query := config.DB.Preload("Profile").Preload("Subscriptions.Plan")

	if search != "" {
		query = query.Where("email LIKE ? OR first_name LIKE ? OR last_name LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	var users []models.User
	var total int64

	query.Model(&models.User{}).Count(&total)
	query.Offset(offset).Limit(limit).Find(&users)

	utils.SuccessResponse(c, "Users retrieved successfully", gin.H{
		"users": users,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func (ac *AdminController) GetUser(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := config.DB.Preload("Profile").Preload("Subscriptions.Plan").Preload("Payments").First(&user, userID).Error; err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	utils.SuccessResponse(c, "User retrieved successfully", user)
}

func (ac *AdminController) BlockUser(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	user.IsBlocked = true
	if err := config.DB.Save(&user).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to block user", err)
		return
	}

	utils.SuccessResponse(c, "User blocked successfully", user)
}

func (ac *AdminController) UnblockUser(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	user.IsBlocked = false
	if err := config.DB.Save(&user).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to unblock user", err)
		return
	}

	utils.SuccessResponse(c, "User unblocked successfully", user)
}

func (ac *AdminController) GetCredKeys(c *gin.Context) {
	var keys []models.CredKey
	if err := config.DB.Preload("User.Profile").Order("created_at DESC").Find(&keys).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve credential keys", err)
		return
	}

	utils.SuccessResponse(c, "Credential keys retrieved successfully", keys)
}

func (ac *AdminController) CreateCredKey(c *gin.Context) {
	var req struct {
		KeyType     string     `json:"key_type" binding:"required"`
		Description string     `json:"description"`
		MaxUses     int        `json:"max_uses"`
		ExpiresAt   *time.Time `json:"expires_at"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	credKey := models.CredKey{
		KeyValue:    utils.GenerateCredKey(req.KeyType),
		KeyType:     req.KeyType,
		Description: req.Description,
		MaxUses:     req.MaxUses,
		ExpiresAt:   req.ExpiresAt,
		IsActive:    true,
	}

	if err := config.DB.Create(&credKey).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create credential key", err)
		return
	}

	utils.SuccessResponse(c, "Credential key created successfully", credKey)
}

func (ac *AdminController) DeleteCredKey(c *gin.Context) {
	keyID := c.Param("id")

	if err := config.DB.Delete(&models.CredKey{}, keyID).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete credential key", err)
		return
	}

	utils.SuccessResponse(c, "Credential key deleted successfully", nil)
}

func (ac *AdminController) GetSubscriptions(c *gin.Context) {
	var subscriptions []models.SubscriptionKey
	if err := config.DB.Preload("User.Profile").Preload("Plan").Order("created_at DESC").Find(&subscriptions).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve subscriptions", err)
		return
	}

	utils.SuccessResponse(c, "Subscriptions retrieved successfully", subscriptions)
}

func (ac *AdminController) GetPayments(c *gin.Context) {
	var payments []models.Payment
	if err := config.DB.Preload("User.Profile").Preload("Subscription.Plan").Order("created_at DESC").Find(&payments).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve payments", err)
		return
	}

	utils.SuccessResponse(c, "Payments retrieved successfully", payments)
}

func (ac *AdminController) GetPlans(c *gin.Context) {
	var plans []models.Plan
	if err := config.DB.Find(&plans).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve plans", err)
		return
	}

	utils.SuccessResponse(c, "Plans retrieved successfully", plans)
}

func (ac *AdminController) CreatePlan(c *gin.Context) {
	var req struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
		Currency    string  `json:"currency"`
		Interval    string  `json:"interval"`
		Features    string  `json:"features"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	plan := models.Plan{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Currency:    req.Currency,
		Interval:    req.Interval,
		Features:    req.Features,
		IsActive:    true,
	}

	if err := config.DB.Create(&plan).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create plan", err)
		return
	}

	utils.SuccessResponse(c, "Plan created successfully", plan)
}

func (ac *AdminController) UpdatePlan(c *gin.Context) {
	planID := c.Param("id")

	var plan models.Plan
	if err := config.DB.First(&plan, planID).Error; err != nil {
		utils.NotFoundResponse(c, "Plan not found")
		return
	}

	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Currency    string  `json:"currency"`
		Interval    string  `json:"interval"`
		Features    string  `json:"features"`
		IsActive    *bool   `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	if req.Name != "" {
		plan.Name = req.Name
	}
	if req.Description != "" {
		plan.Description = req.Description
	}
	if req.Price > 0 {
		plan.Price = req.Price
	}
	if req.Currency != "" {
		plan.Currency = req.Currency
	}
	if req.Interval != "" {
		plan.Interval = req.Interval
	}
	if req.Features != "" {
		plan.Features = req.Features
	}
	if req.IsActive != nil {
		plan.IsActive = *req.IsActive
	}

	if err := config.DB.Save(&plan).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update plan", err)
		return
	}

	utils.SuccessResponse(c, "Plan updated successfully", plan)
}

func (ac *AdminController) DeletePlan(c *gin.Context) {
	planID := c.Param("id")

	if err := config.DB.Delete(&models.Plan{}, planID).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete plan", err)
		return
	}

	utils.SuccessResponse(c, "Plan deleted successfully", nil)
}

func (ac *AdminController) GetSupportTickets(c *gin.Context) {
	var tickets []models.SupportTicket
	if err := config.DB.Preload("User.Profile").Preload("Admin.Profile").Preload("Messages.User.Profile").Order("created_at DESC").Find(&tickets).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve support tickets", err)
		return
	}

	utils.SuccessResponse(c, "Support tickets retrieved successfully", tickets)
}
