package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mauFade/journaly/internal/infrastucture/database"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	} else {
		return defaultValue
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file.")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	db, err := database.ConnectDB(connStr)

	if err != nil {
		log.Fatal("Error connecting to DB: " + err.Error())
	}

}
