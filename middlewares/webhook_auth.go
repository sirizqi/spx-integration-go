package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"time"

	"spx-integration/config"
	"spx-integration/utils"

	"github.com/gin-gonic/gin"
)

func WebhookAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		appID := config.Cfg.AppID
		checkSign := c.GetHeader("check-sign")
		tsStr := c.GetHeader("timestamp")
		rnStr := c.GetHeader("random-num")
		if checkSign == "" || tsStr == "" || rnStr == "" || appID == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ts, _ := strconv.ParseInt(tsStr, 10, 64)
		rn, _ := strconv.ParseInt(rnStr, 10, 64)

		// replay guard (basic)
		if time.Since(time.Unix(ts, 0)) > 10*time.Minute {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// read body
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))

		appIDu64, _ := strconv.ParseUint(appID, 10, 64)
		if !utils.VerifyCheckSign(appIDu64, config.Cfg.AppSecret, ts, rn, bodyBytes, checkSign) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
