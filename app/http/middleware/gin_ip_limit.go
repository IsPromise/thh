package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/leancodebox/goose/serverinfo"
	"net/http"
	"thh/app/bundles/logging"
)

func IpLimit(c *gin.Context) {
	clientIP := c.ClientIP()
	ip, _ := serverinfo.GetLocalIp()
	if len(ip) != 0 && clientIP != ip && clientIP != "::1" {
		logging.Error("clientIp:"+clientIP, " localIp:"+ip)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}
	c.Next()
}
