package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"subscription-saas-backend/config"
	"subscription-saas-backend/models"
	"subscription-saas-backend/utils"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
)

type RazorpayVerifyRequest struct {
	OrderID   string `json:"order_id" binding:"required"`
	PaymentID string `json:"payment_id" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type RazorpayVerifyResponse struct {
	Valid bool   `json:"valid"`
	Msg   string `json:"msg"`
}

// POST /api/v1/payments/razorpay/verify
func (rc *RazorpayController) VerifyPayment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req struct {
		OrderID   string  `json:"order_id" binding:"required"`
		PaymentID string  `json:"payment_id" binding:"required"`
		Signature string  `json:"signature" binding:"required"`
		PlanID    uint    `json:"plan_id" binding:"required"`
		Amount    float64 `json:"amount" binding:"required"`
		Currency  string  `json:"currency" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Generate expected signature
	data := req.OrderID + "|" + req.PaymentID
	h := hmac.New(sha256.New, []byte(config.RazorpayKeySecret))
	h.Write([]byte(data))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	if !hmac.Equal([]byte(expectedSignature), []byte(req.Signature)) {
		c.JSON(http.StatusOK, RazorpayVerifyResponse{Valid: false, Msg: "Signature mismatch"})
		return
	}

	// 1. Validate plan
	var plan models.Plan
	if err := config.DB.First(&plan, req.PlanID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}

	// 2. Create subscription
	now := time.Now()
	var endDate *time.Time
	switch plan.Interval {
	case "monthly":
		t := now.AddDate(0, 1, 0)
		endDate = &t
	case "yearly":
		t := now.AddDate(1, 0, 0)
		endDate = &t
	case "weekly":
		t := now.AddDate(0, 0, 7)
		endDate = &t
	default:
		t := now.AddDate(0, 1, 0)
		endDate = &t
	}
	subscription := models.Subscription{
		UserID:    userID.(uint),
		PlanID:    req.PlanID,
		Status:    "active",
		AutoRenew: true,
		StartDate: &now,
		EndDate:   endDate,
	}
	if err := config.DB.Create(&subscription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription"})
		return
	}

	// 3. Store payment record
	payment := models.Payment{
		UserID:         userID.(uint),
		SubscriptionID: &subscription.ID,
		Amount:         req.Amount,
		Currency:       req.Currency,
		Status:         "completed",
		PaymentMethod:  "Razorpay",
		TransactionID:  req.PaymentID,
		Description:    "Subscription payment for " + plan.Name,
	}
	if err := config.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment record"})
		return
	}

	// 4. Generate subscription key
	key := utils.GenerateSubscriptionKey(userID.(uint))
	subscriptionKey := models.SubscriptionKey{
		OriginalKeyID:    req.PlanID,
		DummyKey:         key,
		AssignedToUserID: new(uint),
		IsUsed:           false,
		AssignedAt:       &now,
	}
	*subscriptionKey.AssignedToUserID = userID.(uint)
	if err := config.DB.Create(&subscriptionKey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription key"})
		return
	}

	// 5. Sync to Firestore and handle error
	firestoreErr := utils.SetUserSubscription(
		fmt.Sprintf("%d", userID.(uint)),
		100000,
		subscriptionKey.DummyKey,
	)

	// 6. Respond with all details, including Firestore error if any
	c.JSON(http.StatusOK, gin.H{
		"valid":            true,
		"msg":              "Signature verified, payment and subscription created",
		"subscription":     subscription,
		"payment":          payment,
		"subscription_key": subscriptionKey,
		"key":              subscriptionKey.DummyKey,
		"firestore_error":  firestoreErr,
	})
}
