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

	// ----------------------------------------------------------
	// 1. Force Service Account Credential for Datastore
	// ----------------------------------------------------------
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:/secrets/serviceAccount.json")

	// (Optional) Debug: Print active service account
	/*
		creds, _ := google.FindDefaultCredentials(context.Background())
		log.Println("ACTIVE SERVICE ACCOUNT:", creds)
	*/

	// ----------------------------------------------------------
	// 2. Load .env file
	// ----------------------------------------------------------
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env variables")
	}

	// ----------------------------------------------------------
	// 3. Connect Database
	// ----------------------------------------------------------
	config.ConnectDatabase()

	// ----------------------------------------------------------
	// 4. Load Razorpay Config
	// ----------------------------------------------------------
	config.LoadRazorpayConfig()

	// ----------------------------------------------------------
	// 5. Auto-Migrate all DB tables
	// ----------------------------------------------------------
	if err := autoMigrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// ----------------------------------------------------------
	// 6. Initialize Google Datastore with CORRECT Project ID
	// ----------------------------------------------------------
	if err := utils.InitDatastore("hive-five-475221"); err != nil {
		log.Fatal("Failed to initialize Datastore:", err)
	}

	// ----------------------------------------------------------
	// 7. Create Gin app
	// ----------------------------------------------------------
	router := gin.Default()

	// ----------------------------------------------------------
	// 8. CORS settings
	// ----------------------------------------------------------
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080", "http://localhost:5173", "http://127.0.0.1:3000", "http://127.0.0.1:8080", "http://127.0.0.1:5173", "http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "x-rtb-fingerprint-id"},
		AllowCredentials: true,
	}))

	// ----------------------------------------------------------
	// 9. Setup routes
	// ----------------------------------------------------------
	routes.SetupRoutes(router)

	// ----------------------------------------------------------
	// 10. Start Server
	// ----------------------------------------------------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}

// Auto-migration
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
