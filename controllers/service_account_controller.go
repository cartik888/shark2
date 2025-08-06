package controllers

import (
	"encoding/json"
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"

	"github.com/gin-gonic/gin"
)

type ServiceAccountController struct{}

func (sac *ServiceAccountController) GetAllEncrypted(c *gin.Context) {
	var accounts []models.ServiceAccount
	if err := config.DB.Find(&accounts).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve service accounts", err)
		return
	}
	encryptedList := []string{}
	for _, acc := range accounts {
		jsonBytes, err := json.Marshal(acc)
		if err != nil {
			utils.InternalServerErrorResponse(c, "JSON marshal failed", err)
			return
		}
		encrypted, err := utils.Encrypt(string(jsonBytes))
		if err != nil {
			utils.InternalServerErrorResponse(c, "Encryption failed", err)
			return
		}
		encryptedList = append(encryptedList, encrypted)
	}
	utils.SuccessResponse(c, "Service accounts (encrypted) retrieved successfully", encryptedList)
}

func (sac *ServiceAccountController) Store(c *gin.Context) {
	var req models.ServiceAccount
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}
	encryptedKey, err := utils.Encrypt(req.PrivateKey)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Encryption failed", err)
		return
	}
	req.PrivateKey = encryptedKey
	if err := config.DB.Create(&req).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to store service account", err)
		return
	}
	utils.SuccessResponse(c, "Service account stored successfully", req)
}

func (sac *ServiceAccountController) GetAll(c *gin.Context) {
	var accounts []models.ServiceAccount
	if err := config.DB.Find(&accounts).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve service accounts", err)
		return
	}
	utils.SuccessResponse(c, "Service accounts retrieved successfully", accounts)
}
