package config

import (
	"fmt"
	"log"
	"os"
)

func LoadDBURI() string {
	
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" || dbHost == "" {
		log.Fatalf("Missing required environment variables")
	}
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		dbUser,
		dbPassword,
		dbName,
		dbPort,
		dbHost,
	)

}