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

	// create content
	data := &goconfluence.Content{
		Type:  "page",           // can also be blogpost
		Title: "Some-Test-Page", // page title
		Ancestors: []goconfluence.Ancestor{
			goconfluence.Ancestor{
				ID: "123456", // ancestor-id optional if you want to create sub-pages
			},
		},
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          "#api-test\nnew sub\npage", // your page content here
				Representation: "storage",
			},
		},
		Version: goconfluence.Version{
			Number: 1,
		},
		Space: goconfluence.Space{
			Key: "SomeSpaceKey", // Space
		},
	}

	c, err := api.CreateContent(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)

	// update content
	data := &goconfluence.Content{
		ID:    "1234567",
		Type:  "page",
		Title: "updated-title",
		Ancestors: []goconfluence.Ancestor{
			goconfluence.Ancestor{
				ID: "2345678",
			},
		},
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          "#api-page\nnew\ncontent",
				Representation: "storage",
			},
		},
		Version: goconfluence.Version{
			Number: 2,
		},
		Space: goconfluence.Space{
			Key: "SomeSpaceKey",
		},
	}

	//c, err := api.CreateContent(data)
	c, err := api.UpdateContent(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)

}
