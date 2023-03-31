package ginLowerControllers

import (
	"github.com/spf13/cast"
	"net/http"
	"strings"
	"thh/arms"

	"github.com/gin-gonic/gin"
)

const configData = `port: 7890
socks-port: 7891
redir-port: 7892
allow-lan: false
mode: global
log-level: silent
external-controller: '0.0.0.0:9090'
proxies:
    - {name: Charles, type: http, server: #{host}, port: 8888}
proxy-groups:
    - {name: global, type: relay, proxies: [Charles]}
`

func GinGetClashConfig(c *gin.Context) {
	host := c.Request.URL.Query().Get("host")
	if len(host) == 0 {
		r := strings.Split(c.Request.Host, ":")
		host = r[0]
	}
	configRep := strings.ReplaceAll(configData, "#{host}", host)
	c.String(http.StatusOK, configRep)
}

func GinGetClashConfigPlus(c *gin.Context) {
	data, _ := arms.FileGetContents("/Users/thh/workspace/thh/storage/logs/config.yaml")
	if len(data) == 0 {
		data = []byte("空")
	}
	c.String(http.StatusOK, cast.ToString(data))
}
