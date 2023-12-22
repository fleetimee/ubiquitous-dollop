package helpers

import (
	"encoding/json"
	"io"
	"os"
)

type Client struct {
	SecretKey string `json:"secret_key"`
}

// Function to read the secret key from client.json
func GetSecretKey() (string, error) {
	jsonFile, err := os.Open("client.json")
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return "", err
	}

	var client Client
	json.Unmarshal(byteValue, &client)

	return client.SecretKey, nil
}
