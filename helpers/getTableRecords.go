package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Record struct {
    ID string `json:"id"`
}

type Records struct {
    Records []Record `json:"records"`
}

func GetTableRecords() {
    apiKey := "Bearer "+ os.Getenv("AIRTABLE_TOKEN")
    baseID := os.Getenv("AIRTABLE_BASEID")
    tableName := os.Getenv("AIRTABLE_TABLE_NAME")
    url := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", baseID, tableName)

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
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

    var records Records
    err = json.NewDecoder(resp.Body).Decode(&records)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

		for _, s := range records.Records {
			fmt.Println(s.ID)
			DeleteTableRecord(s.ID)
		}
}
