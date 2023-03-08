package helpers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type RowData struct {
	Company 		string
	Title 			string
	Date 			string
	URL 			string
	Description 	string
	Employees 		string
}

func AddTableRecords(){
    // Open the CSV file for reading
	file, err := os.Open("searchResult.csv")
	if err != nil {
			fmt.Println(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the CSV data into a slice of structs
	rows := make([]RowData, 0)
	for {
        record, err := reader.Read()
        if err == io.EOF {
                break
        }
        if err != nil {
                panic(err)
        }
        row := RowData{record[0], record[1], record[2], record[3], record[4], record[5]}
        rows = append(rows, row)
	}

	// Create a map to store the values of the column to check for duplicates

	// Iterate over the rows and delete any row that has a duplicate value in column 1
	for i := 1; i < len(rows); i++ {
        num, err := strconv.Atoi(rows[i].Employees)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println(rows[i].Company, rows[i].Title, rows[i].Date, rows[i].URL, rows[i].Description, num);
        addTableRecord(rows[i].Company, rows[i].Title, rows[i].Date, rows[i].URL, rows[i].Description, num)
    }
}

func addTableRecord(company string, title string, date string, layoffURL string, description string, employees int) {
    apiKey := "Bearer "+ os.Getenv("AIRTABLE_TOKEN")
    baseID := os.Getenv("AIRTABLE_BASEID")
    tableName := os.Getenv("AIRTABLE_TABLE_NAME")
    url := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", baseID, tableName)

    // Airtable API key for authentication

    // JSON data representing the record to add
    data := map[string]interface{}{
        "fields": map[string]interface{}{
            "Company": company,
            "Title": title,
            "Date": date,
            "URL": layoffURL,
            "Description": description,
            "Employees": employees,
        },
    }

    // Convert the data to JSON format
    payload, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Create a new HTTP request with the API endpoint URL and payload
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Add the API key to the HTTP request headers for authentication
    req.Header.Set("Authorization", apiKey)
    req.Header.Set("Content-Type", "application/json")

    // Create an HTTP client and send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Check the response status code
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error:", resp.Status)
        return
    }

    // Print the response body
    defer resp.Body.Close()
    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(result)
}
