package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TrackingWebhook(c *gin.Context) { // Webhook Tracking
	var body map[string]any
	_ = c.ShouldBindJSON(&body)
	// TODO: persist to DB / publish to queue
	c.Status(http.StatusOK)
}

func CreateProgressWebhook(c *gin.Context) { // Webhook Create Progress
	var body map[string]any
	_ = c.ShouldBindJSON(&body)
	c.Status(http.StatusOK)
}

func CreateFeedbackWebhook(c *gin.Context) { // Webhook Create Feedback
	var body map[string]any
	_ = c.ShouldBindJSON(&body)
	c.Status(http.StatusOK)
}

func ShippingFeeWebhook(c *gin.Context) { // Webhook Shipping Fee
	var body map[string]any
	_ = c.ShouldBindJSON(&body)
	c.Status(http.StatusOK)
}

func EvidenceProofWebhook(c *gin.Context) { // Webhook Evidence Proof
	var body map[string]any
	_ = c.ShouldBindJSON(&body)
	c.Status(http.StatusOK)
}
