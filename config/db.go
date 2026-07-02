package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (DB *gorm.DB) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_HOST")
	timeZone := "Europe/Kyiv"

	/**
	 * Setup default values for local development
	 */

	if host == "" {
		host = "localhost"
	}
	if user == "" {
		user = "testDbUser"
	}
	if password == "" {
		password = "testDbPassword"
	}
	if dbName == "" {
		dbName = "stockpileManagement"
	}
	if port == "" {
		port = "5432"
	}
	if sslMode == "" {
		sslMode = "allow"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbName, port, sslMode, timeZone)

	/**
	 * Configure GORM logger
	 */

	newLogger := logger.New(
		log.New(
			os.Stdout, "[API] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()

	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(15 * time.Minute)
	}

	log.Println("Database connection established successfully.")

	/**
	 * Auto migrate tables
	 */
	log.Println("Running database migrations...")
	err = db.AutoMigrate()
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	log.Println("Database migrations completed.")

	DB = db
	return
}
