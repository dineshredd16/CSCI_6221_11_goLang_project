package controllers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func ScrapeData (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func main() {
	url := "https://www.cnbc.com/2023/01/20/google-to-lay-off-12000-people-memo-from-ceo-sundar-pichai-says.html?&qsearchterm=google%20layoff"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("error fetching the URL: ", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("the status code is not 200: ", err)
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("error loading the HTML: ", err)
		return
	}

	// Find the header text
	header := doc.Find("header h1").Text()
	if header == "" {
		log.Fatal("headers not found on this page")
		return
	}

	// Extract the company name and job cuts
	var companyName string
	var jobCuts int
	if strings.Contains(header, "Google") {
		companyName = "Google"
		if strings.Contains(header, "lay off") {
			// Extract the number of job cuts using a regular expression
			re := regexp.MustCompile("([0-9]+),([0-9]+)")
			n := re.FindString(header)
			fmt.Println("Match:", n)
			n = strings.Replace(n, ",", "", -1)
			fmt.Sscanf(n, "%d", &jobCuts)
		}
	} else {
		fmt.Println("Company not found on the page")
		return
	}

	// Print the results
	fmt.Println("Company Name:", companyName)
	fmt.Println("Number of Job Cuts:", jobCuts)
}