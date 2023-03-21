package middleware

import (
	"github.com/gin-gonic/gin"
	"thh/arms"
)

func TraceInit(context *gin.Context) {
	arms.MyTraceInit()
	defer arms.MyTraceClean()
	context.Header("X-Powered-By", "PHP/6.0.6")
	context.Next()
}
