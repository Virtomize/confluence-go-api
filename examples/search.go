package main

import (
	"fmt"
	"log"

	"github.com/virtomize/confluence-go-api"
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

}
