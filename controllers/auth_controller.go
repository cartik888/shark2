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
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Bio        string `json:"bio"`
	JobTitle   string `json:"job_title"`
	Company    string `json:"company"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	ZipCode    string `json:"zip_code"`
	Website    string `json:"website"`
	Resume     string `json:"resume"`
	Avatar     string `json:"avatar"`
	CoverImage string `json:"cover_image"`
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
		UserID:     user.ID,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Bio:        req.Bio,
		JobTitle:   req.JobTitle,
		Company:    req.Company,
		Phone:      req.Phone,
		Address:    req.Address,
		City:       req.City,
		State:      req.State,
		Country:    req.Country,
		ZipCode:    req.ZipCode,
		Website:    req.Website,
		Resume:     req.Resume,
		Avatar:     req.Avatar,
		CoverImage: req.CoverImage,
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

type ResetPasswordRequest struct {
	Email           string `json:"email" binding:"required,email"`
	OTP             string `json:"otp" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6"`
}

// POST /api/v1/auth/reset-password
func (ac *AuthController) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	if req.NewPassword != req.ConfirmPassword {
		utils.ValidationErrorResponse(c, "Passwords do not match")
		return
	}

	// Verify OTP
	if !models.VerifyOTP(req.Email, req.OTP) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired OTP", nil)
		return
	}

	// Find user
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to hash password", err)
		return
	}

	user.Password = string(hashedPassword)
	if err := config.DB.Save(&user).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update password", err)
		return
	}

	// Mark OTP as used after successful password reset
	models.MarkOTPUsed(req.Email, req.OTP)
	utils.SuccessResponse(c, "Password reset successful", nil)
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
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Bio        string `json:"bio"`
		JobTitle   string `json:"job_title"`
		Company    string `json:"company"`
		Phone      string `json:"phone"`
		Address    string `json:"address"`
		City       string `json:"city"`
		State      string `json:"state"`
		Country    string `json:"country"`
		ZipCode    string `json:"zip_code"`
		Website    string `json:"website"`
		Resume     string `json:"resume"`
		Avatar     string `json:"avatar"`
		CoverImage string `json:"cover_image"`
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
	profile.Bio = req.Bio
	profile.JobTitle = req.JobTitle
	profile.Company = req.Company
	profile.Phone = req.Phone
	profile.Address = req.Address
	profile.City = req.City
	profile.State = req.State
	profile.Country = req.Country
	profile.ZipCode = req.ZipCode
	profile.Website = req.Website
	profile.Resume = req.Resume
	profile.Avatar = req.Avatar
	profile.CoverImage = req.CoverImage

	if err := config.DB.Save(&profile).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update profile", err)
		return
	}

	utils.SuccessResponse(c, "Profile updated successfully", profile)
}
