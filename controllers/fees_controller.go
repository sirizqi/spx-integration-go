package controllers

import (
	"net/http"
	"spx-integration/models"
	"spx-integration/services"

	"github.com/gin-gonic/gin"
)

func CheckFee(c *gin.Context) {
	var req models.CheckFeeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := services.CheckFee(c, req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetASF(c *gin.Context) {
	var req struct {
		TrackingNos []string `json:"tracking_no_list" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := services.GetASF(c, req.TrackingNos)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
