package middleware

import (
	"fmt"
	"thh/arms/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	logger.Std().WithFields(logrus.Fields{
		"http_status": statusCode,
		"total_time":  latencyTime,
		"ip":          clientIP,
		"method":      reqMethod,
		"uri":         reqUri,
	}).Info("access")

}
