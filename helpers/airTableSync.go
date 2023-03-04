package helpers

import (
    "encoding/csv"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "os"
)

func AirTableSync() {
    // Define your Airtable API key, base ID, and table name
    apiKey := os.Getenv("AIRTABLE_TOKEN")
    baseID := os.Getenv("AIRTABLE_BASEID")
    tableName := os.Getenv("AIRTABLE_TABLE_NAME")

    // Delete all existing records in the table
    deleteRecords(apiKey, baseID, tableName)

    // Open the CSV file and read the contents
    file, err := os.Open("search_results.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // Iterate over each record and add it to the Airtable table
    for _, record := range records {
        addRecord(apiKey, baseID, tableName, record)
    }
}

func deleteRecords(apiKey, baseID, tableName string) {
    // Define the Airtable API endpoint for the table
    endpoint := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", baseID, url.PathEscape(tableName))

    // Create a new HTTP client with the API key as the authorization header
    client := &http.Client{}
    req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

    // Send the HTTP request to delete all records in the table
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    // Check if the response was successful
    if resp.StatusCode != http.StatusOK {
        log.Fatalf("Failed to delete records: %s", resp.Status)
    }

    fmt.Println("All existing records deleted")
}

func addRecord(apiKey, baseID, tableName string, record []string) {
    // Define the Airtable API endpoint for the table
    endpoint := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", baseID, url.PathEscape(tableName))

    // Create a new HTTP client with the API key as the authorization header
    client := &http.Client{}
    data := url.Values{}
    data.Set("fields[Company]", record[0])
    data.Set("fields[Title]", record[1])
    data.Set("fields[Date]", record[2])
    data.Set("fields[URL]", record[3])
    data.Set("fields[Description]", record[4])
    data.Set("fields[Employees]", record[5])

    // Create a new HTTP POST request with the record data as the body
    req, err := http.NewRequest(http.MethodPost, endpoint, nil)
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")

    // Send the HTTP request to add the record to the table
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    // Check if the response was successful
    if resp.StatusCode != http.StatusOK {
        log.Fatalf("Failed to add record: %s", resp.Status)
    }

    fmt.Printf("Record added: %v\n", record)
}
