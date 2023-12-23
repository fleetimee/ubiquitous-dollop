package lib

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("../static/jwt.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger = log.New(logFile, "jwt: ", log.LstdFlags)
}

func GenerateJWT(reqKey string) (string, time.Time, time.Time, error) {
	creationTime := time.Now()

	expirationTimeEnv := os.Getenv("JWT_EXPIRATION_TIME")
	expirationTimeInSeconds, err := strconv.Atoi(expirationTimeEnv)
	if err != nil {
		return "", time.Time{}, time.Time{}, err
	}

	expirationTime := creationTime.Add(time.Duration(expirationTimeInSeconds) * time.Second)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"app_id": reqKey,
		"iss":    "service-fleetime",
		"sub":    "service-fleetime",
		"iat":    creationTime.Unix(),
		"exp":    expirationTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))

	if err != nil {
		return "", time.Time{}, time.Time{}, err
	}

	// Log the creation time, expiration time, and app id
	logData := map[string]interface{}{
		"creation_time":   creationTime,
		"expiration_time": expirationTime,
		"app_id":          reqKey,
	}
	logDataJson, err := json.Marshal(logData)
	if err != nil {
		return "", time.Time{}, time.Time{}, err
	}
	logger.Println(string(logDataJson))

	return tokenString, creationTime, expirationTime, nil
}
