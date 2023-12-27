package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("../static/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	logger = log.New(logFile, "", log.LstdFlags)
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mySigningKey := os.Getenv("JWT_SIGNING_KEY")

		// Extract the JWT token from the request header
		authHeader := c.GetHeader("Authorization")
		splitToken := strings.Split(authHeader, " ") // Split on space

		// Check the length of the splitToken slice
		if len(splitToken) != 2 {
			logger.Println("Error: Authorization header is not correctly formatted")
			c.Next()
			return
		}

		tokenString := splitToken[1] // Take the second part

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the token here
			return []byte(mySigningKey), nil
		})

		if err != nil {
			logger.Println("Error parsing token:", err)
			c.Next()
			return
		}

		// Extract the app_id from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logger.Println("Error extracting claims from token")
			c.Next()
			return
		}

		appId := claims["app_id"]

		// Log the app_id and the accessed endpoint
		logger.Println("App ID:", appId, "Accessed Endpoint:", c.Request.URL.Path)

		// Call the next middleware function
		c.Next()
	}
}
