package controllers

import (
	"net/http"
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.UnauthorizedResponse(c, "Invalid credentials")
		return
	}

	// Check if user is blocked
	if user.IsBlocked {
		utils.ForbiddenResponse(c, "Account is blocked")
		return
	}

	// Check if user is active
	if !user.IsActive {
		utils.ForbiddenResponse(c, "Account is inactive")
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.UnauthorizedResponse(c, "Invalid credentials")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role, user.IsAdmin)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to generate token", err)
		return
	}

	// Load user profile
	config.DB.Preload("Profile").First(&user, user.ID)

	utils.SuccessResponse(c, "Login successful", AuthResponse{
		Token: token,
		User:  user,
	})
}

func (ac *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "User already exists", nil)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to hash password", err)
		return
	}

	// Create user
	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
		IsActive: true,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create user", err)
		return
	}

	// Create user profile
	profile := models.UserProfile{
		UserID:    user.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Company:   req.Company,
	}

	if err := config.DB.Create(&profile).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create user profile", err)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role, user.IsAdmin)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to generate token", err)
		return
	}

	// Load user with profile
	config.DB.Preload("Profile").First(&user, user.ID)

	utils.SuccessResponse(c, "Registration successful", AuthResponse{
		Token: token,
		User:  user,
	})
}

func (ac *AuthController) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user models.User
	if err := config.DB.Preload("Profile").First(&user, userID).Error; err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	utils.SuccessResponse(c, "Profile retrieved successfully", user)
}

func (ac *AuthController) UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Company   string `json:"company"`
		Phone     string `json:"phone"`
		Address   string `json:"address"`
		City      string `json:"city"`
		State     string `json:"state"`
		Country   string `json:"country"`
		ZipCode   string `json:"zip_code"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	var profile models.UserProfile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		utils.NotFoundResponse(c, "Profile not found")
		return
	}

	// Update profile
	profile.FirstName = req.FirstName
	profile.LastName = req.LastName
	profile.Company = req.Company
	profile.Phone = req.Phone
	profile.Address = req.Address
	profile.City = req.City
	profile.State = req.State
	profile.Country = req.Country
	profile.ZipCode = req.ZipCode

	if err := config.DB.Save(&profile).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update profile", err)
		return
	}

	utils.SuccessResponse(c, "Profile updated successfully", profile)
}
