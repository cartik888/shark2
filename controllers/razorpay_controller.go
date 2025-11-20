package controllers

import (
	"net/http"
	"subscription-saas-backend/config"

	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
)

type RazorpayController struct{}

// POST /api/v1/payments/razorpay/order
func (rc *RazorpayController) CreateOrder(c *gin.Context) {
	var req struct {
		Amount   float64 `json:"amount" binding:"required"`
		Currency string  `json:"currency" binding:"required"`
		Receipt  string  `json:"receipt"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	client := razorpay.NewClient(config.RazorpayKeyID, config.RazorpayKeySecret)
	orderAmount := int(req.Amount * 100) // Razorpay expects amount in paise
	orderData := map[string]interface{}{
		"amount":   orderAmount,
		"currency": req.Currency,
		"receipt":  req.Receipt,
	}
	order, err := client.Order.Create(orderData, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order": order})
}
