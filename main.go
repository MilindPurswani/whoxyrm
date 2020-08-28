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
	"strings"
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

// KeywordStruct results more than 50,000 domains
type KeywordStruct struct {
	Status            int               `json:"status"`
	APIQuery          string            `json:"api_query"`
	SearchIdentifier2 SearchIdentifier2 `json:"search_identifier"`
	TotalResults      int               `json:"total_results"`
	TotalPages        int               `json:"total_pages"`
	CurrentPage       int               `json:"current_page"`
	DomainNames       string            `json:"domain_names"`
	APIExecutionTime  float64           `json:"api_execution_time"`
}

// SearchIdentifier2 is of type SearchIdentifier2
type SearchIdentifier2 struct {
	Keyword string `json:"keyword"`
}

func apiGenerator(cn string, name string, email string, keyword string) (string, string) {
	apikey, present := os.LookupEnv("WHOXY_API_KEY")
	mode := "micro"
	if present != true {
		log.Fatal("Are you sure you've set 'WHOXY_API_KEY' environmental variable?")
	}
	var whoxyURL = "http://api.whoxy.com/?key=" + apikey + "&reverse=whois"
	if len(cn) > 0 {
		whoxyURL = whoxyURL + "&mode=micro&company=" + url.QueryEscape(cn)
	} else if len(name) > 0 {
		whoxyURL = whoxyURL + "&mode=micro&name=" + url.QueryEscape(name)
	} else if len(email) > 0 {
		whoxyURL = whoxyURL + "&mode=micro&email=" + url.QueryEscape(email)
	} else if len(keyword) > 0 {
		whoxyURL = whoxyURL + "&mode=domains&keyword=" + url.QueryEscape(keyword)
		mode = "domain"
	}
	return whoxyURL, mode
}

func getResult(page int, url string, mode string) int {

	var whoxyURL = url + "&page=" + strconv.Itoa(page)
	// var whoxyURL = "https://www.whoxy.com/sample/reverseWhoisMicro.json"
	resp, err := http.Get(whoxyURL)
	if err != nil {
		log.Fatal(err)
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if mode == "micro" {
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
	} else {
		respHandler := KeywordStruct{}
		respString := string(respData)
		err = json.Unmarshal([]byte(respString), &respHandler)
		if err != nil {
			log.Fatal(err)
		}
		d := strings.Split(respHandler.DomainNames, ",")
		for j := 0; j < len(d); j++ {
			fmt.Println(d[j])
		}
		return respHandler.TotalPages
	}
}

func main() {
	var cn string
	flag.StringVar(&cn, "company-name", "", "Company Name for which you need to get all the assets from whoxy")
	var name string
	flag.StringVar(&name, "name", "", "Domain Name for which you need to get all the assets from whoxy")
	var email string
	flag.StringVar(&email, "email", "", "Email address for which you need to get all the assets from whoxy")
	var keyword string
	flag.StringVar(&keyword, "keyword", "", "Keyword for which you need to get all the assets from whoxy. Returns much more results but high chances of false positives. Get 50,000 results in one request.")
	var rCount int
	flag.IntVar(&rCount, "result-count", -1, "The count of results that you need to fetch from the API. Keep in mind that 1 request will give 2500 domains only. So, if you want to fetch 10,000 results, the tool will make 4 requests. Make sure you have sufficient credits available.")
	flag.Parse()
	if flag.NArg() > 0 {
		log.Fatal("Kindly check the docs whoxy -h for usage")
	}

	whoxyURL, mode := apiGenerator(cn, name, email, keyword)
	pageCount := getResult(rCount, whoxyURL, mode)
	if pageCount > 1 {
		for i := 1; i <= pageCount; i++ {
			_ = getResult(i, whoxyURL, mode)
		}
	}

}
