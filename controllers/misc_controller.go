package controllers

import (
	"net/http"
	"spx-integration/services"

	"github.com/gin-gonic/gin"
)

func GetAddressLink(c *gin.Context) {
	resp, err := services.GetAddressLink(c)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
