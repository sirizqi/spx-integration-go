package routes

import (
	"spx-integration/controllers"
	"spx-integration/middlewares"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	// Orders
	r.GET("/pickup-time", controllers.GetPickupTime)  // 1.1
	r.POST("/orders", controllers.CreateOrder)        // 1.2
	r.POST("/orders/track", controllers.TrackOrder)   // 1.3
	r.POST("/orders/cancel", controllers.CancelOrder) // 1.4
	r.POST("/orders/awb", controllers.GetAWB)         // 1.5

	// Fees
	r.POST("/fees/check", controllers.CheckFee) // 2.1
	r.POST("/fees/asf", controllers.GetASF)     // 2.2

	// Misc
	r.POST("/address/download-url", controllers.GetAddressLink) // 5.1

	// Webhooks (protected)
	w := r.Group("/webhooks", middlewares.WebhookAuth())
	{
		w.POST("/tracking", controllers.TrackingWebhook)             // 4.1
		w.POST("/order-progress", controllers.CreateProgressWebhook) // 4.2
		w.POST("/order-feedback", controllers.CreateFeedbackWebhook) // 4.3
		w.POST("/shipping-fee", controllers.ShippingFeeWebhook)      // 4.6
		w.POST("/evidence", controllers.EvidenceProofWebhook)        // 4.7
	}

	return r
}
