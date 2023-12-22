package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var DBPostgres *gorm.DB

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

	// Logging to a file.
	f, _ := os.Create("../static/sqlserver.log")
	newLogger := logger.New(
		log.New(f, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	// Open connection to database
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Print something if connection success
	println("Connection to database success")

	DB = db
}

func ConnectPostgres() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("failed to load env file")
	}

	// Mapping variable from .env file
	host := os.Getenv("PG_DB_HOST")
	port := os.Getenv("PG_DB_PORT")
	user := os.Getenv("PG_DB_USER")
	pass := os.Getenv("PG_DB_PASSWORD")
	dbname := os.Getenv("PG_DB_NAME")

	// Create connection string postgres
	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"

	// Logging to a file.
	f, _ := os.Create("../static/postgres.log")
	newLogger := logger.New(
		log.New(f, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	// Open connection to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Print something if connection success
	println("Connection to database success")

	DBPostgres = db
}
