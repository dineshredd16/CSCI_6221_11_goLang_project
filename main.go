package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const baseURL = "https://api.queryly.com/cnbc/json.aspx"

type SearchResult struct {
	Results []struct {
		Title string `json:"cn:title"`
		Date  string `json:"_pubDate"`
		Url   string `json:"cn:liveURL"`
		Desc  string `json:"description"`
	} `json:"results"`
}

func main() {
	// List of company names to search for
	companies := []string{"Google", "Apple", "Meta", "Microsoft"}

	// Similar words to include in the search query
	similarWords := []string{"layoff"}

	// Open CSV file for writing
	file, err := os.Create("search_results.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row to CSV file
	writer.Write([]string{"Company", "Title", "Date", "URL", "Description"})

	// Loop through companies and search for "layoff" and company name on CNBC website
	for _, company := range companies {
		// Build query parameters
		queryParams := url.Values{
			"queryly_key":       {"31a35d40a9a64ab3"},
			"query":             {fmt.Sprintf("%s %s", strings.Join(similarWords, " "), company)},
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
			"fromdate":          {"-12m"},
		}

		// Make request to CNBC API
		for page := 0; page < 11; page++ {
			fmt.Printf("Extracting Page# %d for %s\n", page+1, company)
			queryParams.Set("endindex", fmt.Sprintf("%d", page*100))
			resp, err := http.Get(baseURL + "?" + queryParams.Encode())
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			// Parse response into SearchResult struct
			var result SearchResult
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				fmt.Println(err)
				return
			}
			/*
				// Write each result to CSV file
				for _, item := range result.Results {
					writer.Write([]string{company, item.Title, item.Date, item.Url, item.Desc})
				}
			*/
			// Write each result to CSV file
			for _, item := range result.Results {
				// Extract the number of people affected from the description
				var numPeople string
				numPeopleRegex := regexp.MustCompile(`[0-9]+([,.][0-9]+)*(\s)?(thousand|million|billion)?(\s)?(people|employees|workers|staff)`)
				match := numPeopleRegex.FindStringSubmatch(item.Desc)
				if len(match) > 0 {
					numPeople = match[0]
				}

				// Exclude percentage numbers from the description
				descWithoutPercentages := strings.ReplaceAll(item.Desc, "%", "")
				descWithoutPercentages = strings.TrimSpace(descWithoutPercentages)

				writer.Write([]string{company, item.Title, item.Date, item.Url, descWithoutPercentages, numPeople})
			}

		}
	}

	fmt.Println("Search results saved to search_results.csv")
}
