package main

import (
	"fmt"
	"log"

	goconfluence "github.com/virtomize/confluence-go-api"
)

func main() {
	api, err := goconfluence.NewAPI("https://<your-domain>.atlassian.net/wiki/rest/api", "<username>", "<api-token>")
	if err != nil {
		log.Fatal(err)
	}

	// define your Search parameters
	query := goconfluence.SearchQuery{
		CQL: "space=SomeSpace",
	}

	// execute search
	result, err := api.Search(query)
	if err != nil {
		log.Fatal(err)
	}

	// loop over results
	for _, v := range result.Results {
		fmt.Printf("%+v\n", v)
	}

	// search example with paging using SearchWithNext and Links.Next
	next := ""
	for {
		resp, err := api.SearchWithNext(query, next)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range result.Results {
			fmt.Printf("%+v\n", v)
		}
		next = resp.Links.Next
		if next == "" {
			break
		}
		log.Printf("Using next page: %s", next)
	}
}
