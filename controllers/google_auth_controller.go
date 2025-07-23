package controllers

import (
	"context"
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

type GoogleAuthRequest struct {
	IdToken string `json:"id_token" binding:"required"`
}

type GoogleUserInfo struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
}

type GoogleAuthController struct{}

func (gc *GoogleAuthController) GoogleLogin(c *gin.Context) {
	var req GoogleAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Invalid request data")
		return
	}

	payload, err := idtoken.Validate(context.Background(), req.IdToken, "55308792729-o7ifngca382s8rels0d1iusmuuklj9sb.apps.googleusercontent.com")
	if err != nil {
		utils.UnauthorizedResponse(c, "Invalid Google token")
		return
	}

	claims := payload.Claims
	email, _ := claims["email"].(string)
	emailVerified, _ := claims["email_verified"].(bool)
	picture, _ := claims["picture"].(string)
	givenName, _ := claims["given_name"].(string)
	familyName, _ := claims["family_name"].(string)

	if !emailVerified {
		utils.UnauthorizedResponse(c, "Email not verified by Google")
		return
	}

	var user models.User
	err = config.DB.Where("email = ?", email).Preload("Profile").First(&user).Error
	if err != nil {
		// User does not exist, create new user and profile
		user = models.User{
			Email:    email,
			Role:     "user",
			IsActive: true,
		}
		if err := config.DB.Create(&user).Error; err != nil {
			utils.InternalServerErrorResponse(c, "Failed to create user", err)
			return
		}
		profile := models.UserProfile{
			UserID:    user.ID,
			FirstName: givenName,
			LastName:  familyName,
			Avatar:    picture,
		}
		config.DB.Create(&profile)
	} else {
		// User exists, update profile if missing or empty
		var updated bool
		if user.Profile == nil {
			profile := models.UserProfile{
				UserID:    user.ID,
				FirstName: givenName,
				LastName:  familyName,
				Avatar:    picture,
			}
			config.DB.Create(&profile)
			updated = true
		} else {
			// Update fields if empty
			p := user.Profile
			if p.FirstName == "" && givenName != "" {
				p.FirstName = givenName
				updated = true
			}
			if p.LastName == "" && familyName != "" {
				p.LastName = familyName
				updated = true
			}
			if p.Avatar == "" && picture != "" {
				p.Avatar = picture
				updated = true
			}
			if updated {
				config.DB.Save(p)
			}
		}
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role, user.IsAdmin)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to generate token", err)
		return
	}

	config.DB.Preload("Profile").First(&user, user.ID)

	utils.SuccessResponse(c, "Google login successful", AuthResponse{
		Token: token,
		User:  user,
	})
}
