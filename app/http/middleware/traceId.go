package middleware

import (
	"github.com/gin-gonic/gin"
)

func TraceInit(context *gin.Context) {
	//luckrand.MyTraceInit()
	//defer luckrand.MyTraceClean()
	context.Header("X-Powered-By", "PHP/6.0.6")
	context.Next()
}
