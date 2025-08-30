package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// We just ack 200 fast, then you can enqueue to worker/message bus if needed.

func TrackingWebhook(c *gin.Context) { // 4.1
	var body map[string]any
	_ = c.ShouldBindJSON(&body)
	// TODO: persist to DB / publish to queue
	c.Status(http.StatusOK)
}

func CreateProgressWebhook(c *gin.Context) { // 4.2
	var body map[string]any
	_ = c.ShouldBindJSON(&body)
	c.Status(http.StatusOK)
}

func CreateFeedbackWebhook(c *gin.Context) { // 4.3
	var body map[string]any
	_ = c.ShouldBindJSON(&body)
	c.Status(http.StatusOK)
}

func ShippingFeeWebhook(c *gin.Context) { // 4.6
	var body map[string]any
	_ = c.ShouldBindJSON(&body)
	c.Status(http.StatusOK)
}

func EvidenceProofWebhook(c *gin.Context) { // 4.7
	var body map[string]any
	_ = c.ShouldBindJSON(&body)
	c.Status(http.StatusOK)
}
