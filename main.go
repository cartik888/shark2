package main

import (
	"log"
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	config.ConnectDatabase()

	// Auto migrate models
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
		&models.Plan{},
		&models.Subscription{},
		&models.SubscriptionEvent{},
		&models.Payment{},
		&models.PaymentMethod{},
		&models.Invoice{},
		&models.PasswordReset{},
		&models.CredKey{},
		&models.SubscriptionKey{},
		&models.SupportTicket{},
		&models.SupportMessage{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Seed initial data
	seedInitialData()

	// Initialize Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetupRoutes(router)

	// Start server
	log.Println("🚀 Server starting on http://localhost:8080")
	log.Println("📊 Admin Panel: http://localhost:3000/admin")
	log.Println("🔑 Default Admin: admin@shark.com / admin123")
	router.Run(":8080")
}

func seedInitialData() {
	// Seed plans
	var planCount int64
	config.DB.Model(&models.Plan{}).Count(&planCount)
	if planCount == 0 {
		plans := []models.Plan{
			{
				Name:        "Basic",
				Description: "Perfect for individuals getting started",
				Price:       9.99,
				Currency:    "USD",
				Interval:    "monthly",
				Features:    "5 Projects, 10GB Storage, Email Support, Basic Analytics",
				IsActive:    true,
			},
			{
				Name:        "Pro",
				Description: "Great for growing businesses",
				Price:       29.99,
				Currency:    "USD",
				Interval:    "monthly",
				Features:    "Unlimited Projects, 100GB Storage, Priority Support, Advanced Analytics, Team Collaboration, API Access",
				IsActive:    true,
			},
			{
				Name:        "Enterprise",
				Description: "For large organizations",
				Price:       99.99,
				Currency:    "USD",
				Interval:    "monthly",
				Features:    "Everything in Pro, Custom Integrations, Dedicated Support, SLA Guarantee, White-label Options",
				IsActive:    true,
			},
		}
		config.DB.Create(&plans)
		log.Println("✅ Sample plans created")
	}

	// Seed credential keys
	var keyCount int64
	config.DB.Model(&models.CredKey{}).Count(&keyCount)
	if keyCount == 0 {
		keys := []models.CredKey{
			{KeyValue: "SHARK-ACTIVATION-2024-001", KeyType: "activation", IsActive: true},
			{KeyValue: "SHARK-ACTIVATION-2024-002", KeyType: "activation", IsActive: true},
			{KeyValue: "SHARK-ACTIVATION-2024-003", KeyType: "activation", IsActive: true},
			{KeyValue: "SHARK-ACTIVATION-2024-004", KeyType: "activation", IsActive: true},
			{KeyValue: "SHARK-ACTIVATION-2024-005", KeyType: "activation", IsActive: true},
			{KeyValue: "SHARK-TRIAL-2024-001", KeyType: "trial", IsActive: true},
			{KeyValue: "SHARK-TRIAL-2024-002", KeyType: "trial", IsActive: true},
			{KeyValue: "SHARK-PARTNER-2024-001", KeyType: "partner", IsActive: true},
		}
		config.DB.Create(&keys)
		log.Println("✅ Sample credential keys created")
	}

	// Seed admin user
	var adminCount int64
	config.DB.Model(&models.User{}).Where("email = ?", "admin@shark.com").Count(&adminCount)
	if adminCount == 0 {
		// Hash password: admin123
		hashedPassword := "$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi"
		
		admin := models.User{
			Email:    "admin@shark.com",
			Password: hashedPassword,
			Role:     "admin",
			IsAdmin:  true,
			IsActive: true,
		}
		config.DB.Create(&admin)

		// Create admin profile
		profile := models.UserProfile{
			UserID:    admin.ID,
			FirstName: "Admin",
			LastName:  "User",
			Company:   "Shark Technologies",
		}
		config.DB.Create(&profile)
		log.Println("✅ Admin user created (admin@shark.com / admin123)")
	}
}
