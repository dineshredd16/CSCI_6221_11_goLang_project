package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const (
	baseURL    = "https://www.cnbc.com/search/"
	queryParam = "Google Layoff"
	searchTerm = "Google Layoff"
	source     = ""
	typeParam  = ""
)

type SearchResult struct {
	Results []struct {
		Title string `json:"cn:title"`
		Date  string `json:"_pubDate"`
		Url   string `json:"cn:liveURL"`
		Desc  string `json:"description"`
	} `json:"results"`
}

func main() {
	baseURL := "https://api.queryly.com/cnbc/json.aspx"
	queryParams := url.Values{
		"queryly_key":       {"31a35d40a9a64ab3"},
		"query":             {"google" + " layoff OR " + "google" + " layoffs"},
		"endindex":          {"0"},
		"batchsize":         {"100"},
		"callback":          {""},
		"showfaceted":       {"true"},
		"timezoneoffset":    {"-120"},
		"facetedfields":     {"formats"},
		"facetedkey":        {"formats|"},
		"facetedvalue":      {"!Press Release|"},
		"needtoptickers":    {"1"},
		"additionalindexes": {"4cd6f71fbf22424d,937d600b0d0d4e23,3bfbe40caee7443e,626fdfcd96444f28"},
	}

	var allin [][]string
	for page := 0; page < 11; page++ {
		fmt.Printf("Extracting Page# %d\n", page+1)
		queryParams.Set("endindex", fmt.Sprintf("%d", page*100))
		resp, err := http.Get(baseURL + "?" + queryParams.Encode())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		var result SearchResult
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, item := range result.Results {
			allin = append(allin, []string{item.Title, item.Date, item.Url, item.Desc})
		}
	}

	// Create a new CSV file and write the data to it
	filePath := "C:/Users/15716/Desktop/6212/search_results.csv"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row
	header := []string{"Title", "Date", "Url", "Description"}
	writer.Write(header)

	// Write the data rows
	for _, row := range allin {
		writer.Write(row)
	}
}
