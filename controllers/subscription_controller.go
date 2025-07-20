package controllers

import (
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"

	"github.com/gin-gonic/gin"
)

type SubscriptionController struct{}

func (sc *SubscriptionController) GetPlans(c *gin.Context) {
	var plans []models.Plan
	if err := config.DB.Where("is_active = ?", true).Find(&plans).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve plans", err)
		return
	}

	utils.SuccessResponse(c, "Plans retrieved successfully", plans)
}

func (sc *SubscriptionController) GetPlan(c *gin.Context) {
	planID := c.Param("id")

	var plan models.Plan
	if err := config.DB.First(&plan, planID).Error; err != nil {
		utils.NotFoundResponse(c, "Plan not found")
		return
	}

	utils.SuccessResponse(c, "Plan retrieved successfully", plan)
}

func (sc *SubscriptionController) ValidateKey(c *gin.Context) {
	var req struct {
		KeyValue string `json:"key_value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	// Check subscription key first
	var subscriptionKey models.SubscriptionKey
	if err := config.DB.Preload("AssignedToUser").Preload("OriginalKey").Where("dummy_key = ? AND is_used = ?", req.KeyValue, false).First(&subscriptionKey).Error; err == nil {
		utils.SuccessResponse(c, "Subscription key is valid", gin.H{
			"key_type":         "subscription",
			"assigned_to_user": subscriptionKey.AssignedToUser,
			"original_key_id":  subscriptionKey.OriginalKeyID,
			"is_used":          subscriptionKey.IsUsed,
		})
		return
	}

	// Check credential key
	var credKey models.CredKey
	if err := config.DB.Where("key_value = ? AND is_active = ?", req.KeyValue, true).First(&credKey).Error; err != nil {
		utils.NotFoundResponse(c, "Invalid key")
		return
	}

	utils.SuccessResponse(c, "Credential key is valid", gin.H{
		"key_type":     "credential",
		"key":          credKey,
		"current_uses": credKey.CurrentUses,
		"max_uses":     credKey.MaxUses,
	})
}
