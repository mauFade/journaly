package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	journalservice "github.com/mauFade/journaly/internal/application/service/journal-service"
	userservice "github.com/mauFade/journaly/internal/application/service/user-service"
	"github.com/mauFade/journaly/internal/infrastucture/database"
	"github.com/mauFade/journaly/internal/infrastucture/database/repository"
	"github.com/mauFade/journaly/internal/presentation/http/server"
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
	defer db.Close()

	ur := repository.NewUserRepository(db)
	jr := repository.NewJournalRepository(db)

	us := userservice.NewUserService(ur)
	js := journalservice.NewJournalService(jr)

	p := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(us, js, p)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("error starting http server.")
	}
}
