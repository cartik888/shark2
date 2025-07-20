package middlewares

import (
	"strings"
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.UnauthorizedResponse(c, "Authorization header required")
			c.Abort()
			return
		}

		// Check if header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.UnauthorizedResponse(c, "Invalid authorization header format")
			c.Abort()
			return
		}

		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			utils.UnauthorizedResponse(c, "Token required")
			c.Abort()
			return
		}

		// Validate token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			utils.UnauthorizedResponse(c, "Invalid token")
			c.Abort()
			return
		}

		// Check if user still exists and is active
		var user models.User
		if err := config.DB.First(&user, claims.UserID).Error; err != nil {
			utils.UnauthorizedResponse(c, "User not found")
			c.Abort()
			return
		}

		if user.IsBlocked {
			utils.ForbiddenResponse(c, "Account is blocked")
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)
		c.Set("is_admin", claims.IsAdmin)

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("is_admin")
		if !exists || !isAdmin.(bool) {
			utils.ForbiddenResponse(c, "Admin access required")
			c.Abort()
			return
		}

		c.Next()
	}
}
