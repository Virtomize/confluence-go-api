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

	// get content by content id
	c, err := api.GetContent("12345678", goconfluence.ContentQuery{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)
}
