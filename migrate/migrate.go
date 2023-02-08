package main

import (
	"webScraperBackend/initializers"
	"webScraperBackend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Site{})
}