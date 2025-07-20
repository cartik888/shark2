package routes

import (
	"subscription-saas-backend/controllers"
	"subscription-saas-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Initialize controllers
	authController := &controllers.AuthController{}
	userController := &controllers.UserController{}
	subscriptionController := &controllers.SubscriptionController{}
	paymentController := &controllers.PaymentController{}
	keyController := &controllers.KeyController{}
	adminController := &controllers.AdminController{}

	// API v1 routes
	v1 := router.Group("/api/v1")

	// Public routes (no authentication required)
	public := v1.Group("/")
	{
		// Authentication routes
		public.POST("/auth/login", authController.Login)
		public.POST("/auth/register", authController.Register)

		// Public subscription routes
		public.GET("/plans", subscriptionController.GetPlans)
		public.GET("/plans/:id", subscriptionController.GetPlan)

		// Key validation routes (for desktop app)
		public.POST("/keys/validate", subscriptionController.ValidateKey)
		public.POST("/keys/validate-cred", keyController.ValidateCredKey)
		public.POST("/keys/validate-subscription", keyController.ValidateSubscriptionKey)
	}

	// Protected routes (authentication required)
	protected := v1.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		// User profile routes
		protected.GET("/auth/profile", authController.GetProfile)
		protected.GET("/me", authController.GetProfile)
		protected.PUT("/auth/profile", authController.UpdateProfile)

		// User dashboard and data
		protected.GET("/user/dashboard", userController.GetDashboard)
		protected.GET("/user/subscriptions", userController.GetSubscriptions)
		protected.GET("/subscriptions/my", userController.GetSubscriptions)
		protected.GET("/user/payments", userController.GetPayments)
		protected.POST("/user/buy-plan", userController.BuyPlan)
		protected.GET("/user/plan-usage", userController.GetPlanUsage)

		// NEW: Get MY keys after payment
		protected.GET("/user/my-keys", userController.GetMyKeys)
		protected.POST("/user/key-by-payment", userController.GetKeyByPayment)

		// Payment routes
		protected.GET("/payments", paymentController.GetPayments)
		protected.GET("/payments/:id", paymentController.GetPayment)
		protected.POST("/payment-methods", paymentController.CreatePaymentMethod)
		protected.GET("/payment-methods", paymentController.GetPaymentMethods)

		// Checkout routes
		protected.POST("/payments/checkout", paymentController.CreateCheckout)
		protected.POST("/payments/process", paymentController.ProcessCheckout)

		// Subscription Key routes
		protected.GET("/subscription-keys", keyController.GetUserSubscriptionKeys)
		protected.GET("/subscription-keys/active", keyController.GetActiveSubscriptionKeys)
		protected.GET("/subscription-keys/payment/:payment_id", keyController.GetSubscriptionKeysByPayment)
		protected.GET("/subscription-keys/checkout/:checkout_id", keyController.GetSubscriptionKeyByCheckout)
		protected.POST("/subscription-keys/use", keyController.UseSubscriptionKey)
	}

	// Admin routes (admin authentication required)
	admin := v1.Group("/admin")
	admin.Use(middlewares.AuthMiddleware())
	admin.Use(middlewares.AdminMiddleware())
	{
		// Admin dashboard
		admin.GET("/dashboard", adminController.GetDashboard)

		// User management
		admin.GET("/users", adminController.GetUsers)
		admin.GET("/users/:id", adminController.GetUser)
		admin.POST("/users/:id/block", adminController.BlockUser)
		admin.POST("/users/:id/unblock", adminController.UnblockUser)

		// Credential key management
		admin.GET("/cred-keys", adminController.GetCredKeys)
		admin.POST("/cred-keys", adminController.CreateCredKey)
		admin.DELETE("/cred-keys/:id", adminController.DeleteCredKey)

		// Subscription management
		admin.GET("/subscriptions", adminController.GetSubscriptions)

		// Payment management
		admin.GET("/payments", adminController.GetPayments)

		// Plan management
		admin.GET("/plans", adminController.GetPlans)
		admin.POST("/plans", adminController.CreatePlan)
		admin.PUT("/plans/:id", adminController.UpdatePlan)
		admin.DELETE("/plans/:id", adminController.DeletePlan)

		// Support management
		admin.GET("/support", adminController.GetSupportTickets)
	}

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Shark SaaS Backend is running",
		})
	})
}
