package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	fmt.Println("Environment Variables Loaded")
}

func init() {
	LoadEnv()
}
