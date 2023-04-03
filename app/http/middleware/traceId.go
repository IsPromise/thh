package middleware

import (
	"thh/arms"

	"github.com/gin-gonic/gin"
)

func TraceInit(context *gin.Context) {
	arms.MyTraceInit()
	defer arms.MyTraceClean()
	context.Header("X-Powered-By", "PHP/6.0.6")
	context.Next()
}
