package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"service-fleetime/cmd/helpers"
	"service-fleetime/cmd/lib"
	"time"

	"github.com/gin-gonic/gin"
)

type Config map[string]string

func GetToken(c *gin.Context) {
	configFile, err := os.ReadFile("../client.json")
	if err != nil {
		panic(err)
	}

	var config Config
	json.Unmarshal(configFile, &config)

	var reqConfig Config
	if err := c.ShouldBindJSON(&reqConfig); err != nil {

		helpers.APIResponse(c, err.Error(), http.StatusBadRequest, nil)
		return
	}

	// Check if the request body contains exactly one key-value pair
	if len(reqConfig) != 1 {

		helpers.APIResponse(
			c,
			"Request body should contain exactly one key-value pair",
			http.StatusBadRequest,
			nil,
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
				helpers.APIResponse(
					c,
					fmt.Sprintf("Error generating token: %v", err),
					http.StatusInternalServerError,
					nil,
				)

				return
			}

			helpers.APIResponse(c, "Authorized", http.StatusOK, gin.H{
				"token":           tokenString,
				"creation_time":   creationTime.Format(time.RFC3339),
				"expiration_time": expirationTime.Format(time.RFC3339),
			})
			return
		} else {
			helpers.APIResponse(c, "Invalid value for key", http.StatusUnauthorized, nil)
			return
		}
	} else {
		helpers.APIResponse(c, "Invalid key", http.StatusUnauthorized, nil)
		return
	}

}
