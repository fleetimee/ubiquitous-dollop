package main

import (
	"service-fleetime/cmd/routes"
	"service-fleetime/config"
)

func main() {
	config.Connect()
	config.ConnectPostgres()
	routes.FetcherRoutes()
}
