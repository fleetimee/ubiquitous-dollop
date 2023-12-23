package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"service-fleetime/cmd/controller"
	"service-fleetime/cmd/helpers"
	"service-fleetime/cmd/lib"
	"service-fleetime/cmd/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

var NewLogger logger.Interface

func ClearDirectory(dir string) error {
	// Open the directory.
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	// Read all file names from the directory.
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	// Loop over the file names and delete each one.
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}

	return nil
}

type Config map[string]string

func FetcherRoutes() {

	configFile, err := ioutil.ReadFile("../client.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(configFile))

	var config Config
	json.Unmarshal(configFile, &config)

	fmt.Println(config)

	f, _ := os.Create("../static/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	r.StaticFS("/static", http.Dir("../static"))
	r.GET("/clear-static", func(c *gin.Context) {
		// Open the directory.
		d, err := os.Open("../static")
		if err != nil {
			helpers.ErrorResponse(c, err.Error())
			return
		}
		defer d.Close()

		// Read all file names from the directory.
		names, err := d.Readdirnames(-1)
		if err != nil {
			helpers.ErrorResponse(c, err.Error())
			return
		}

		// Loop over the file names and clear each one.
		for _, name := range names {
			f, err := os.OpenFile(filepath.Join("../static", name), os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				helpers.ErrorResponse(c, err.Error())
				return
			}
			defer f.Close()

			err = f.Truncate(0)
			if err != nil {
				helpers.ErrorResponse(c, err.Error())
				return
			}
		}

		helpers.APIResponse(c, "Logs Cleared", http.StatusOK, nil)
	})

	r.POST("/token-client", func(c *gin.Context) {
		var reqConfig Config
		if err := c.ShouldBindJSON(&reqConfig); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the request body contains exactly one key-value pair
		if len(reqConfig) != 1 {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "Request body should contain exactly one key-value pair"},
			)
			return
		}

		// Extract the key and value from the request body
		var reqKey, reqValue string
		for key, value := range reqConfig {
			reqKey = key
			reqValue = value
		}

		// Check if the key-value pair matches a pair in the config map
		if configValue, ok := config[reqKey]; ok {
			if reqValue == configValue {
				// generate and return JWT token
				tokenString, creationTime, expirationTime, err := lib.GenerateJWT(reqKey)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
					return
				}

				// c.JSON(http.StatusOK, gin.H{"token": tokenString})

				helpers.APIResponse(c, "Authorized", http.StatusOK, gin.H{
					"token":           tokenString,
					"creation_time":   creationTime.Format(time.RFC3339),
					"expiration_time": expirationTime.Format(time.RFC3339),
				})
				return
			}

			helpers.APIResponse(c, "Invalid app id or key", http.StatusUnauthorized, nil)
		}

		helpers.APIResponse(c, "Unknown error", http.StatusUnauthorized, nil)
	})

	r.GET("/fetcher", controller.GetData)
	r.GET("/fetcher/all", middleware.AuthMiddleware(), controller.FetchAllEmployee)
	r.GET("/fetcher/all/send", controller.FetchAllEmployeeAndSendToPostgres)

	r.Run(":8080")
}
