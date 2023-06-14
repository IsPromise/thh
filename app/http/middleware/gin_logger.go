package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leancodebox/goose/jsonopt"
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
	requestData, _ := c.Get("requestData")

	info := fmt.Sprintf("access http_status:%v total_time:%v ip:%v method:%v uri:%v body:%s",
		statusCode, latencyTime, clientIP,
		reqMethod, reqUri, jsonopt.Encode(requestData),
	)
	logging.Info(info)

}
