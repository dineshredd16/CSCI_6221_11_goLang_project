package helpers

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Row struct {
	Company 			string
	Title 				string
	Date 					string
	URL 					string
	Description 	string
	Employees 		string
}

func QueryDuplicatesFromCSV() {
	// Open the CSV file for reading
	file, err := os.Open("searchResult.csv")
	if err != nil {
			fmt.Println(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the CSV data into a slice of structs
	rows := make([]Row, 0)
	for {
			record, err := reader.Read()
			if err == io.EOF {
					break
			}
			if err != nil {
					panic(err)
			}
			row := Row{record[0], record[1], record[2], record[3], record[4], record[5]}
			rows = append(rows, row)
	}

	// Create a map to store the values of the column to check for duplicates
	columnMap := make(map[string]bool)

	// Iterate over the rows and delete any row that has a duplicate value in column 1
	for i := 0; i < len(rows); i++ {
			if columnMap[rows[i].Company] && columnMap[rows[i].Employees] {
					rows = append(rows[:i], rows[i+1:]...)
					i--
			} else {
					columnMap[rows[i].Company] = true
					columnMap[rows[i].Employees] = true
			}
	}

	// Open the CSV file for writing
	outputFile, err := os.Create("searchResult.csv")
	if err != nil {
			panic(err)
	}
	defer outputFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(outputFile)

	// Write the updated rows to the CSV file
	for _, row := range rows {
			record := []string{row.Company, row.Title, row.Date, row.URL, row.Description, row.Employees}
			err := writer.Write(record)
			if err != nil {
					panic(err)
			}
	}

	// Flush the CSV writer
	writer.Flush()

	fmt.Println("search results saved to searchResult.csv")
}