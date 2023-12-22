package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"service-fleetime/cmd/controller"
	"service-fleetime/cmd/helpers"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

var mySigningKey = []byte("secret")

func FetcherRoutes() {

	configFile, err := ioutil.ReadFile("../client.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(configFile))

	var config Config
	json.Unmarshal(configFile, &config)

	fmt.Println(config)

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("../static/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

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

	authMiddleware := func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Remove the "Bearer " prefix
		if len(tokenString) > 7 && strings.ToUpper(tokenString[0:7]) == "BEARER " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})

		if token != nil && token.Valid {
			c.Next()
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "That's not even a token"})
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Timing is everything"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Couldn't handle this token: " + err.Error()})
		}
		c.Abort()
	}

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
				expirationTime := time.Now().Add(5 * time.Second).Unix()
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"app_id": reqKey,
					"sub":    "service-fleetime",
					"exp":    expirationTime,
				})
				tokenString, err := token.SignedString(mySigningKey)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
					return
				}

				c.JSON(http.StatusOK, gin.H{"token": tokenString})
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid app id or key"})
	})

	r.GET("/fetcher", controller.GetData)
	r.GET("/fetcher/all", authMiddleware, controller.FetchAllEmployee)
	r.GET("/fetcher/all/send", controller.FetchAllEmployeeAndSendToPostgres)

	r.Run(":8080")
}
