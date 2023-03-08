package helpers

import (
    "fmt"
    "net/http"
    "os"
)

func DeleteTableRecord(recordID string) {
		apiKey := "Bearer "+ os.Getenv("AIRTABLE_TOKEN")
		baseID := os.Getenv("AIRTABLE_BASEID")
		tableName := os.Getenv("AIRTABLE_TABLE_NAME")
    url := fmt.Sprintf("https://api.airtable.com/v0/%s/%s/%s", baseID, tableName, recordID)

    client := &http.Client{}

    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    req.Header.Add("Authorization", apiKey)
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Failed to delete record. Status code: %d", resp.StatusCode)
        os.Exit(1)
    }

    fmt.Println("Record deleted successfully")
}
