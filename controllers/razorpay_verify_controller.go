package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"subscription-saas-backend/config"

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
	var req RazorpayVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Generate expected signature
	data := req.OrderID + "|" + req.PaymentID
	h := hmac.New(sha256.New, []byte(config.RazorpayKeySecret))
	h.Write([]byte(data))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	if hmac.Equal([]byte(expectedSignature), []byte(req.Signature)) {
		c.JSON(http.StatusOK, RazorpayVerifyResponse{Valid: true, Msg: "Signature verified"})
	} else {
		c.JSON(http.StatusOK, RazorpayVerifyResponse{Valid: false, Msg: "Signature mismatch"})
	}
}
