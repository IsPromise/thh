package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thh/bundles/logging"
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

	info := fmt.Sprintf("access http_status:%v total_time:%v ip:%v method:%v uri:%v", statusCode,
		latencyTime,
		clientIP,
		reqMethod,
		reqUri)
	logging.Info(info)

}
