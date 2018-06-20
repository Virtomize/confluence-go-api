package main

import (
	"fmt"
	"log"

	"github.com/cseeger-epages/confluence-go-api"
)

func main() {
	api, err := goconfluence.NewAPI("https://<your-domain>.atlassian.net", "<username>", "<api-token>")
	if err != nil {
		log.Fatal(err)
	}

	// get comments of a specific page
	res, err := api.GetComments("1234567")
	if err != nil {
		log.Fatal(err)
	}

	// loop over results
	for _, v := range res.Results {
		fmt.Printf("%+v\n", v)
	}

	// get attachments of a specific page
	res, err := api.GetAttachments("1234567")
	if err != nil {
		log.Fatal(err)
	}

	// loop over results
	for _, v := range res.Results {
		fmt.Printf("%+v\n", v)
	}

	// get child pages of a specific page
	res, err := api.GetChildPages("1234567")
	if err != nil {
		log.Fatal(err)
	}

	// loop over results
	for _, v := range res.Results {
		fmt.Printf("%+v\n", v)
	}
}
