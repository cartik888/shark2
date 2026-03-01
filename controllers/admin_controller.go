package controllers

import (
	"strconv"
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// ...existing code...

// DeleteUser deletes a user by ID
func (ac *AdminController) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		utils.ValidationErrorResponse(c, "User ID is required")
		return
	}

	// Convert userID to uint
	uid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		utils.ValidationErrorResponse(c, "Invalid user ID")
		return
	}

	// Prevent admin from deleting themselves
	currentUser, exists := c.Get("user")
	if exists {
		if u, ok := currentUser.(*models.User); ok && u != nil && u.ID == uint(uid) {
			utils.ValidationErrorResponse(c, "You cannot delete your own account")
			return
		}
	}

	// Delete user and related profile
	if err := config.DB.Delete(&models.User{}, uid).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete user", err)
		return
	}
	if err := config.DB.Where("user_id = ?", uid).Delete(&models.UserProfile{}).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete user profile", err)
		return
	}

	utils.SuccessResponse(c, "User deleted successfully", nil)
}

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

// GetTotalUsers returns total active platform users
func (ac *AdminController) GetTotalUsers(c *gin.Context) {
	var totalUsers, activeUsers, inactiveUsers int64

	// Get total users
	if err := config.DB.Model(&models.User{}).Count(&totalUsers).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get total users", err)
		return
	}

	// Get active users
	if err := config.DB.Model(&models.User{}).Where("is_active = ? AND is_blocked = ?", true, false).Count(&activeUsers).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get active users", err)
		return
	}

	// Get inactive users
	if err := config.DB.Model(&models.User{}).Where("is_active = ? OR is_blocked = ?", false, true).Count(&inactiveUsers).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get inactive users", err)
		return
	}

	utils.SuccessResponse(c, "Total users retrieved successfully", gin.H{
		"total_users":    totalUsers,
		"active_users":   activeUsers,
		"inactive_users": inactiveUsers,
	})
}

// GetActiveSubscriptions returns active paying customers
func (ac *AdminController) GetActiveSubscriptions(c *gin.Context) {
	var activeSubscriptions, totalSubscriptions int64
	var totalPayingRevenue float64

	// Get total subscriptions
	if err := config.DB.Model(&models.Subscription{}).Count(&totalSubscriptions).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get total subscriptions", err)
		return
	}

	// Get active subscriptions
	if err := config.DB.Model(&models.Subscription{}).Where("status = ?", "active").Count(&activeSubscriptions).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get active subscriptions", err)
		return
	}

	// Get total revenue from active subscriptions
	if err := config.DB.Model(&models.Payment{}).
		Joins("JOIN subscriptions ON payments.subscription_id = subscriptions.id").
		Where("payments.status = ? AND subscriptions.status = ?", "completed", "active").
		Select("COALESCE(SUM(payments.amount), 0)").
		Scan(&totalPayingRevenue).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get revenue", err)
		return
	}

	utils.SuccessResponse(c, "Active subscriptions retrieved successfully", gin.H{
		"active_subscriptions": activeSubscriptions,
		"total_subscriptions":  totalSubscriptions,
		"paying_customers":     activeSubscriptions,
		"total_revenue":        totalPayingRevenue,
	})
}

// GetMonthlyRevenue returns monthly revenue with comparison to last month
func (ac *AdminController) GetMonthlyRevenue(c *gin.Context) {
	currentMonth := time.Now()
	currentMonthStart := time.Date(currentMonth.Year(), currentMonth.Month(), 1, 0, 0, 0, 0, currentMonth.Location())
	currentMonthEnd := currentMonthStart.AddDate(0, 1, 0)

	previousMonthStart := currentMonthStart.AddDate(0, -1, 0)
	previousMonthEnd := currentMonthStart

	var currentMonthRevenue, previousMonthRevenue float64

	// Get current month revenue
	if err := config.DB.Model(&models.Payment{}).
		Where("status = ? AND created_at >= ? AND created_at < ?", "completed", currentMonthStart, currentMonthEnd).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&currentMonthRevenue).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get current month revenue", err)
		return
	}

	// Get previous month revenue
	if err := config.DB.Model(&models.Payment{}).
		Where("status = ? AND created_at >= ? AND created_at < ?", "completed", previousMonthStart, previousMonthEnd).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&previousMonthRevenue).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get previous month revenue", err)
		return
	}

	// Calculate percentage change
	var percentageChange float64
	if previousMonthRevenue > 0 {
		percentageChange = ((currentMonthRevenue - previousMonthRevenue) / previousMonthRevenue) * 100
	} else if currentMonthRevenue > 0 {
		percentageChange = 100
	}

	utils.SuccessResponse(c, "Monthly revenue retrieved successfully", gin.H{
		"current_month":      currentMonthStart.Format("2006-01"),
		"current_revenue":    currentMonthRevenue,
		"previous_month":     previousMonthStart.Format("2006-01"),
		"previous_revenue":   previousMonthRevenue,
		"revenue_difference": currentMonthRevenue - previousMonthRevenue,
		"percentage_change":  percentageChange,
		"trend":              map[bool]string{true: "up", false: "down"}[currentMonthRevenue >= previousMonthRevenue],
	})
}

// GetUserGrowth returns new user registrations grouped by month
func (ac *AdminController) GetUserGrowth(c *gin.Context) {
	// Get the last 12 months of user registrations
	type UserGrowthData struct {
		Month    string
		NewUsers int64
	}

	var growthData []UserGrowthData

	// Query to get user registrations by month for the last 12 months (MySQL compatible)
	if err := config.DB.Model(&models.User{}).
		Where("created_at >= DATE_SUB(NOW(), INTERVAL 12 MONTH)").
		Select("DATE_FORMAT(created_at, '%Y-%m') as month, COUNT(*) as new_users").
		Group("DATE_FORMAT(created_at, '%Y-%m')").
		Order("month ASC").
		Scan(&growthData).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get user growth data", err)
		return
	}

	// Convert to gin.H format
	var formattedGrowthData []gin.H
	for _, item := range growthData {
		formattedGrowthData = append(formattedGrowthData, gin.H{
			"month":     item.Month,
			"new_users": item.NewUsers,
		})
	}

	// Get total user count
	var totalUsers int64
	config.DB.Model(&models.User{}).Count(&totalUsers)

	// Get current month new users
	currentMonthStart := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
	var currentMonthNewUsers int64
	config.DB.Model(&models.User{}).
		Where("created_at >= ?", currentMonthStart).
		Count(&currentMonthNewUsers)

	utils.SuccessResponse(c, "User growth data retrieved successfully", gin.H{
		"total_users":       totalUsers,
		"current_month_new": currentMonthNewUsers,
		"growth_by_month":   formattedGrowthData,
		"last_updated":      time.Now(),
	})
}
