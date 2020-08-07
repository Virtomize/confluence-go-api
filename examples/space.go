package main

import (
	"fmt"
	"log"

	"github.com/xinminlabs/confluence-go-api"
)

func main() {
	api, err := goconfluence.NewAPI("https://<your-domain>.atlassian.net/wiki/rest/api", "<username>", "<api-token>")
	if err != nil {
		log.Fatal(err)
	}

	spaces, err := api.GetAllSpaces(goconfluence.AllSpacesQuery{
		Type:  "global",
		Start: 0,
		Limit: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, space := range spaces.Results {
		fmt.Printf("Space Key: %s\n", space.Key)
	}
}
