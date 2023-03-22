package ginLowerControllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
)

type Router struct {
	EntryPoints []string
	Rule        string
	Service     string
}
type Service struct {
	LoadBalancer struct {
		PassHostHeader bool
		Servers        []map[string]string
	}
}

type buildYml struct {
	Http struct {
		Routers  map[string]Router
		Services map[string]Service
	}
}

func TraefikProvider(ctx *gin.Context) {
	name := "http-service"
	address := "http://host.docker.internal:8001/"
	rule := "PathPrefix(`/xxxx/xxxxi/`)"

	b := buildYml{
		Http: struct {
			Routers  map[string]Router
			Services map[string]Service
		}{Routers: map[string]Router{
			name + "p": {
				EntryPoints: []string{"web"},
				Rule:        rule,
				Service:     name,
			},
		}, Services: map[string]Service{
			name: {
				LoadBalancer: struct {
					PassHostHeader bool
					Servers        []map[string]string
				}{
					PassHostHeader: false,
					Servers: []map[string]string{
						{"url": address},
					},
				},
			},
		}},
	}
	yamlData, _ := yaml.Marshal(b)
	ctx.String(http.StatusOK, cast.ToString(yamlData))
}
