package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("failed to load env file")
	}

	// Mapping variable from .env file
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Create connection string
	dsn := "sqlserver://" + user + ":" + pass + "@" + host + ":" + port + "?database=" + dbname + "&connection+timeout=30"

	// Open connection to database
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Print something if connection success
	println("Connection to database success")

	DB = db
}
