package main

import (
	"webScraperBackend/helpers"
	"webScraperBackend/initializers"
)
func main() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDatabase()
  // List of company names to search for
  companies := helpers.CompanyNames()
  // Similar words to include in the search query
  similarWords := helpers.SearchQueryTerms()
  // Open CSV file for writing and sending results to it
  helpers.StartScraper(companies, similarWords)
}
