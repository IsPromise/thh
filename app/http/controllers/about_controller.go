package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leancodebox/goose/serverinfo"
	"github.com/spf13/cast"
	"net/http"
	"runtime"
	"strings"
	"thh/app/http/controllers/component"
	"thh/bundles/logging"
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
	return component.SuccessResponse(cast.ToString(m.Alloc/1024) + "KB")
}

func About() component.Response {
	return component.SuccessResponse(component.DataMap{
		"message": "Hello~ Now you see a json from gin",
	})
}

func SysInfo() component.Response {
	var s serverinfo.Server
	var err error
	s.Os = serverinfo.InitOS()
	if s.Cpu, err = serverinfo.InitCPU(); err != nil {
		logging.ErrIf(err)

	}

	if s.Ram, err = serverinfo.InitRAM(); err != nil {
		logging.ErrIf(err)
	}

	if s.Disk, err = serverinfo.InitDisk(); err != nil {
		logging.ErrIf(err)
	}

	return component.SuccessResponse(s)
}
