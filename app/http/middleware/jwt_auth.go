package middleware

import (
	"net/http"
	"strings"
	"thh/arms/jwt"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

// JWTAuth4Fiber
// 如果未获取到 x-token 则非法登陆
// 如果已经过期 则推出
// 如果只是则在header返回新的token
func JWTAuth4Fiber() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var token string
		token = c.Get("x-token")
		if token == "" {
			return c.JSON("未登陆")
		}
		userId, newToken, err := jwt.VerifyTokenWithFresh(token)
		if err != nil {
			errorMsg := err.Error()
			if err == jwt.TokenExpired {
				errorMsg = "授权已过期"
			}
			return c.JSON(errorMsg)
		}
		if token != newToken {
			c.Set("new-token", newToken)
		}
		c.Locals("userId", userId)
		return c.Next()
	}
}

func JWTAuth4Gin(c *gin.Context) {
	var token string
	//token = c.GetHeader("x-token")
	token = c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	if token == "" {
		c.JSON(http.StatusUnauthorized, "未登陆")
		c.Abort()
		return
	}
	userId, newToken, err := jwt.VerifyTokenWithFresh(token)
	if err != nil {
		errorMsg := err.Error()
		if err == jwt.TokenExpired {
			errorMsg = "授权已过期"
		}
		c.JSON(http.StatusUnauthorized, errorMsg)
	}
	if token != newToken {
		c.Header("new-token", newToken)
	}
	c.Set("userId", userId)
	c.Next()
}
