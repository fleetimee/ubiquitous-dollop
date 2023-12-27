package client

import (
	"service-fleetime/cmd/controller"
	"service-fleetime/cmd/middleware"

	"github.com/gin-gonic/gin"
)

func FetchEmployee(r *gin.Engine) {
	r.GET("/employees", middleware.AuthMiddleware(), controller.FetchAllEmployee)
}
