package helpers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	// "time"
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

func buildResults(companies []string, similarWords []string, writer *csv.Writer) {
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

				count := 0
				for _, s := range companies {
					if (strings.Contains(strings.ToLower(item.Title), strings.ToLower(s))) && (strings.Contains(strings.ToLower(item.Title), strings.ToLower(company))) {
						count++
					}
				}

				if (len(numPeople) > 0) && (count == 1){
                    reg, err := regexp.Compile("[^0-9]+")
                    if err != nil {
                        fmt.Println(err)
                    }
                    processedString := reg.ReplaceAllString(numPeople, "")
					writer.Write([]string{company, item.Title, item.Date, item.Url, descWithoutPercentages, processedString})
				}
			}
		}
	}
}

func buildCSV(companies []string, similarWords []string){
	file, err := os.Create("searchResult.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Write header row to CSV file
	writer.Write([]string{"Company", "Title", "Date", "URL", "Description", "Employees"})
	buildResults(companies, similarWords, writer)
	defer file.Close()
	writer.Flush()
	QueryDuplicatesFromCSV()
}

func StartScraper(companies []string, similarWords []string) {
  // Create a ticker that ticks every 1 minute
  // ticker := time.NewTicker(1 * time.Second)

  // Loop through the ticker and execute the function
  // for range ticker.C {
    buildCSV(companies, similarWords)
  // }
}