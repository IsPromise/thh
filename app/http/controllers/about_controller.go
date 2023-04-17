package controllers

import (
	"net/http"
	"runtime"
	"strings"
	"thh/app/http/controllers/component"
	"thh/arms"
	"thh/arms/logger"

	"github.com/gin-gonic/gin"

	"github.com/spf13/cast"
)

func Api(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"msg": "OK",
	})
}

const (
	contentTypeHTML      = "text/html"
	errorCodeNotFound    = 404
	errorMessageNotFound = "路由未定义，请确认 url 和请求方法是否正确。"
)

func NotFound(c *gin.Context) {
	acceptString := c.GetHeader("Accept")
	if strings.Contains(acceptString, contentTypeHTML) {
		c.Redirect(http.StatusTemporaryRedirect, "/actor")
		return
	}
	c.JSON(http.StatusNotFound, component.DataMap{
		"error_code":    errorCodeNotFound,
		"error_message": errorMessageNotFound,
	})
}

func GetUseMem() component.Response {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return component.SuccessResponse(cast.ToString(m.Alloc/1024/8) + "kb")
}

func About() component.Response {
	return component.SuccessResponse(component.DataMap{
		"message": "Hello~ Now you see a json from gin",
	})
}

func SysInfo() component.Response {
	var s arms.Server
	var err error
	s.Os = arms.InitOS()
	if s.Cpu, err = arms.InitCPU(); err != nil {
		logger.ErrIf(err)

	}

	if s.Ram, err = arms.InitRAM(); err != nil {
		logger.ErrIf(err)
	}

	if s.Disk, err = arms.InitDisk(); err != nil {
		logger.ErrIf(err)
	}

	return component.SuccessResponse(s)
}
