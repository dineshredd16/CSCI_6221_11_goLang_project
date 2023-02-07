package initializers

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	databaseURL := os.Getenv("WEB_SCRAPER_DEV_DATABASE")
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	if err != nil {
		log.Fatal("unbale to connect to the database error: ", err)
	}
}