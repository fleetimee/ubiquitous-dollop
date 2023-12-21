package routes

import (
	"service-fleetime/cmd/controller"

	"github.com/gin-gonic/gin"
)

func FetcherRoutes() {
	r := gin.Default()

	r.GET("/fetcher", controller.GetData)

	r.Run(":8080")
}
