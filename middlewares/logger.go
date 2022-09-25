package middlewares

import (
	"github.com/gin-gonic/gin"
	lib "octaviusfarrel.dev/latihan_web/lib/log"
)

type LoggerMiddleware struct{}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (LoggerMiddleware) GetLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := lib.NewLogger()
		log.Log(map[string]interface{}{
			"handler_location": ctx.HandlerName(),
		}, "Request started", 0)
		ctx.Next()
		log.Log(map[string]interface{}{
			"handler_location": ctx.HandlerName(),
		}, "Request finished", 0)
	}
}
