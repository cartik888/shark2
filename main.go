package main

import (
	"log"
	"os"
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/routes"
	"subscription-saas-backend/utils"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file (this is OK if running in production with real env vars)")
	}
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to database
	config.ConnectDB()

	// Auto-migrate database tables
	if err := autoMigrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize Firestore
	if err := utils.InitFirestore("hive-five-465116"); err != nil {
		log.Fatal("Failed to initialize Firestore:", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			allowed := []string{
				"http://localhost:3000",
				"http://localhost:8080",
				"http://localhost:5173/",
				"http://127.0.0.1:3000",
				"http://127.0.0.1:8080",
				"http://127.0.0.1:5173",
			}
			for _, o := range allowed {
				if o == origin {
					return true
				}
			}
			return false
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetupRoutes(router)

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}

func autoMigrate() error {
	return config.DB.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
		&models.PasswordReset{},
		&models.Plan{},
		&models.Subscription{},
		&models.Payment{},
		&models.PaymentMethod{},
		&models.Invoice{},
		&models.CredKey{},
		&models.SubscriptionKey{},
		&models.SupportTicket{},
		&models.SupportMessage{},
	)
}
