package controllers

import (
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type KeyController struct{}

type ValidateKeyRequest struct {
	KeyValue string `json:"key_value" binding:"required"`
}

func (kc *KeyController) ValidateCredKey(c *gin.Context) {
	var req ValidateKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	var credKey models.CredKey
	if err := config.DB.Where("key_value = ? AND is_active = ?", req.KeyValue, true).First(&credKey).Error; err != nil {
		utils.ValidationErrorResponse(c, "Invalid or inactive key")
		return
	}

	// Check if key has reached max uses
	if credKey.MaxUses > 0 && credKey.CurrentUses >= credKey.MaxUses {
		utils.ValidationErrorResponse(c, "Key has reached maximum usage limit")
		return
	}

	utils.SuccessResponse(c, "Key is valid", gin.H{
		"key_type":     credKey.KeyType,
		"description":  credKey.Description,
		"expires_at":   credKey.ExpiresAt,
		"max_uses":     credKey.MaxUses,
		"current_uses": credKey.CurrentUses,
	})
}

func (kc *KeyController) UseCredKey(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req ValidateKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	var credKey models.CredKey
	if err := config.DB.Where("key_value = ? AND is_active = ?", req.KeyValue, true).First(&credKey).Error; err != nil {
		utils.ValidationErrorResponse(c, "Invalid or inactive key")
		return
	}

	// Check if key has reached max uses
	if credKey.MaxUses > 0 && credKey.CurrentUses >= credKey.MaxUses {
		utils.ValidationErrorResponse(c, "Key has reached maximum usage limit")
		return
	}

	// Check if user already used this key
	if credKey.UsedBy != nil && *credKey.UsedBy == userID {
		utils.ValidationErrorResponse(c, "You have already used this key")
		return
	}

	// Generate subscription key
	subscriptionKeyValue := utils.GenerateSubscriptionKey(credKey.ID)
	now := time.Now()

	// Create subscription key
	subscriptionKey := models.SubscriptionKey{
		OriginalKeyID:    credKey.ID,
		DummyKey:         subscriptionKeyValue,
		AssignedToUserID: &userID,
		IsUsed:           false,
		AssignedAt:       &now,
	}

	if err := config.DB.Create(&subscriptionKey).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create subscription key", err)
		return
	}

	// Update credential key usage
	credKey.CurrentUses++
	if credKey.MaxUses == 1 {
		credKey.UsedBy = &userID
		credKey.UsedAt = &now
	}
	config.DB.Save(&credKey)

	utils.SuccessResponse(c, "Subscription key generated successfully", gin.H{
		"subscription_key": subscriptionKeyValue,
		"key_type":         credKey.KeyType,
		"description":      credKey.Description,
	})
}

func (kc *KeyController) GetUserSubscriptionKeys(c *gin.Context) {
	userID := c.GetUint("user_id")

	var subscriptionKeys []models.SubscriptionKey
	if err := config.DB.Preload("OriginalKey").Preload("AssignedToUser").Where("assigned_to_user_id = ?", userID).Find(&subscriptionKeys).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch subscription keys", err)
		return
	}

	utils.SuccessResponse(c, "Subscription keys retrieved successfully", subscriptionKeys)
}

func (kc *KeyController) ValidateSubscriptionKey(c *gin.Context) {
	var req ValidateKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	var subscriptionKey models.SubscriptionKey
	if err := config.DB.Preload("OriginalKey").Preload("AssignedToUser").Where("dummy_key = ?", req.KeyValue).First(&subscriptionKey).Error; err != nil {
		utils.ValidationErrorResponse(c, "Invalid subscription key")
		return
	}

	// Check if key is already used
	if subscriptionKey.IsUsed {
		utils.ValidationErrorResponse(c, "Subscription key has already been used")
		return
	}

	utils.SuccessResponse(c, "Subscription key is valid", gin.H{
		"key_type":    subscriptionKey.OriginalKey.KeyType,
		"description": subscriptionKey.OriginalKey.Description,
		"assigned_to": subscriptionKey.AssignedToUser.Email,
		"assigned_at": subscriptionKey.AssignedAt,
	})
}

func (kc *KeyController) UseSubscriptionKey(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req ValidateKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	var subscriptionKey models.SubscriptionKey
	if err := config.DB.Preload("OriginalKey").Where("dummy_key = ? AND assigned_to_user_id = ?", req.KeyValue, userID).First(&subscriptionKey).Error; err != nil {
		utils.ValidationErrorResponse(c, "Invalid subscription key or not assigned to you")
		return
	}

	// Check if key is already used
	if subscriptionKey.IsUsed {
		utils.ValidationErrorResponse(c, "Subscription key has already been used")
		return
	}

	// Mark subscription key as used
	subscriptionKey.IsUsed = true
	if err := config.DB.Save(&subscriptionKey).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update subscription key", err)
		return
	}

	utils.SuccessResponse(c, "Subscription key used successfully", gin.H{
		"key_type":    subscriptionKey.OriginalKey.KeyType,
		"description": subscriptionKey.OriginalKey.Description,
		"used_at":     time.Now(),
	})
}
