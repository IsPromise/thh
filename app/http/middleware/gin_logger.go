package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thh/app/bundles/logging"
	"time"
)

func GinLogger(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	endTime := time.Now()
	latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))
	reqMethod := c.Request.Method
	reqUri := c.Request.RequestURI
	statusCode := c.Writer.Status()
	clientIP := c.ClientIP()
	requestData, _ := c.Get("requestData")
	logging.Info("access",
		"http_status", statusCode,
		"total_time", latencyTime,
		"ip", clientIP,
		"ip", reqMethod,
		"uri", reqUri,
		"body", requestData,
	)
}
