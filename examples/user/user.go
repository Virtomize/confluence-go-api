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

	// get current user information
	currentUser, err := api.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", currentUser)

	// get anonymous user information
	anonUser, err := api.AnonymousUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", anonUser)

	// get user by username or accountId
	user, err := api.User("someuser")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", user)
}
