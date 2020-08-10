package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// ResponseStruct returns response_struct
type ResponseStruct struct {
	Status           int              `json:"status"`
	APIQuery         string           `json:"api_query"`
	SearchIdentifier SearchIdentifier `json:"search_identifier"`
	TotalResults     int              `json:"total_results"`
	TotalPages       int              `json:"total_pages"`
	CurrentPage      int              `json:"current_page"`
	SearchResult     []SearchResult   `json:"search_result"`
	APIExecutionTime float64          `json:"api_execution_time"`
}

// SearchIdentifier is of the type search_identifier
type SearchIdentifier struct {
	Name string `json:"name"`
}

// SearchResult is of the type search_result
type SearchResult struct {
	Num           int    `json:"num"`
	DomainName    string `json:"domain_name"`
	QueryTime     string `json:"query_time"`
	CreateDate    string `json:"create_date"`
	UpdateDate    string `json:"update_date"`
	ExpiryDate    string `json:"expiry_date"`
	RegistrarName string `json:"registrar_name"`
}

func getResult(page int, cn string, apikey string) int {
	var whoxyURL = "http://api.whoxy.com/?key=" + apikey + "&reverse=whois&mode=micro&company=" + url.QueryEscape(cn) + "&page=" + strconv.Itoa(page)
	// fmt.Println(whoxyURL2)
	// var whoxyURL = "https://www.whoxy.com/sample/reverseWhoisMicro.json"
	resp, err := http.Get(whoxyURL)
	if err != nil {
		log.Fatal(err)
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	respHandler := ResponseStruct{}
	respString := string(respData)
	err = json.Unmarshal([]byte(respString), &respHandler)
	if err != nil {
		log.Fatal(err)
	}
	for j := 0; j < len(respHandler.SearchResult); j++ {
		fmt.Println(respHandler.SearchResult[j].DomainName)
	}
	return respHandler.TotalPages
}

func main() {
	var cn string
	flag.StringVar(&cn, "company-name", "", "Company Name for which you need to get all the assets from whoxy")
	var rCount int
	flag.IntVar(&rCount, "result-count", -1, "The count of results that you need to fetch from the API. Keep in mind that 1 request will give 50000 domains only. So, if you want to fetch 100,000 results, the tool will make 2 requests. Make sure you have sufficient credits available.")
	flag.Parse()
	if flag.NArg() > 0 {
		log.Fatal("Kindly check the docs whoxy -h for usage")
	}
	apikey, present := os.LookupEnv("WHOXY_API_KEY")
	if present != true {
		log.Fatal("Are you sure you've set 'WHOXY_API_KEY' environmental variable?")
	}
	pageCount := getResult(rCount, cn, apikey)
	if pageCount > 1 {
		for i := 1; i <= pageCount; i++ {
			_ = getResult(i, cn, apikey)
		}
	}

}
