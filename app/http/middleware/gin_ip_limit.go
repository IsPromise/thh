package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/leancodebox/goose/serverinfo"
	"net/http"
)

func IpLimit(c *gin.Context) {
	clientIP := c.ClientIP()
	ip, _ := serverinfo.GetLocalIp()
	if len(ip) != 0 && clientIP != ip && clientIP != "::1" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}
	c.Next()
}
