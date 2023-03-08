package main

import (
	"webScraperBackend/helpers"
	"webScraperBackend/initializers"
)
func main() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDatabase()
  companies := helpers.CompanyNames()
  similarWords := helpers.SearchQueryTerms()
  helpers.StartScraper(companies, similarWords)
  helpers.GetTableRecords()
  helpers.AddTableRecords()
}
