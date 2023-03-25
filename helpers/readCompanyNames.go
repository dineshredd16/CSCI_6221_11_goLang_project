package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readCompanyNamesFromCSV() []string{
	// Open the CSV file
	file, err := os.Open("companies.csv")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Extract the company names and store them in an array
	var companies []string
	for _, record := range records {
		companies = append(companies, record[0])
	}

	// Print the array of company names
	fmt.Println(companies)
	return companies
}