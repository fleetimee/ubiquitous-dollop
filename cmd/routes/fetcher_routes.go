package routes

import (
	"io"
	"os"
	"service-fleetime/cmd/routes/client"
	notfound "service-fleetime/cmd/routes/not-found"
	"service-fleetime/cmd/routes/utilities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

var NewLogger logger.Interface

func FetcherRoutes() {

	f, _ := os.Create("../static/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// gin.SetMode("release")

	r := gin.Default()

	// NO ROUTES FOUND
	notfound.Roadblock(r)

	// SERVER ROUTES
	utilities.ServeStatic(r)
	utilities.ClearLog(r)

	// AUTH ROUTES
	client.FetchToken(r)

	// EMPLOYEE ROUTES
	client.FetchEmployee(r)

	serverPort := os.Getenv("SERVER_PORT")

	if serverPort == "" {
		serverPort = "8080"
	}

	r.Run(":" + serverPort)
}
