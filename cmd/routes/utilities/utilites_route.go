package utilities

import (
	"service-fleetime/cmd/controller"
	"service-fleetime/cmd/middleware"

	"github.com/gin-gonic/gin"
)

func ServeStatic(r *gin.Engine) {
	r.GET("/static/*filepath", middleware.ServerMiddleware(), controller.ServeStatic)
}

func ClearLog(r *gin.Engine) {
	r.GET("/clear-static", middleware.ServerMiddleware(), controller.ClearLog)
}
