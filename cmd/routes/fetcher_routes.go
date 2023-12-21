package routes

import (
	"service-fleetime/cmd/controller"

	"github.com/gin-gonic/gin"
)

func FetcherRoutes() {
	r := gin.Default()

	r.GET("/fetcher", controller.GetData)
	r.GET("/fetcher/all", controller.FetchAllEmployee)
	r.GET("/fetcher/all/send", controller.FetchAllEmployeeAndSendToPostgres)

	r.Run(":8080")
}
