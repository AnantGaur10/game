package db

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"game/models"
	"game/pkg/config"
)

var (
	Db   *gorm.DB
	once sync.Once
)

// Setup database connection
func setupDatabase() {

	// Build the connection string
	dsn := config.LoadDBURI()
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	err = Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v\n", err)
	}
	log.Println("Database Connected")
}

func init() {
	once.Do(setupDatabase)
}
