package middleware

import (
	"errors"
	"net/http"
	"os"
	"service-fleetime/cmd/helpers"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if len(tokenString) > 7 && strings.ToUpper(tokenString[0:7]) == "BEARER " {
			tokenString = tokenString[7:]
		}

		mySigningKey := os.Getenv("JWT_SIGNING_KEY")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		})

		if token != nil && token.Valid {
			c.Next()
		} else if errors.Is(err, jwt.ErrTokenMalformed) {

			helpers.APIResponse(c, "That's not even a token", http.StatusUnauthorized, nil)
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {

			helpers.APIResponse(c, "Invalid signature", http.StatusUnauthorized, nil)
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			helpers.APIResponse(c, "Token Expired", http.StatusUnauthorized, nil)
		} else {

			helpers.APIResponse(c, "Couldn't handle this token: "+err.Error(), http.StatusUnauthorized, nil)
		}
		c.Abort()
	}
}
