/*
Package goconfluence implementing atlassian's Confluence API

Simple example:

	//Initialize a new API instance
	api, err := goconfluence.NewAPI(
		"https://<your-domain>.atlassian.net/wiki/rest/api",
		"<username>",
		"<api-token>",
	)
	if err != nil {
		log.Fatal(err)
	}

	// get current user information
	currentUser, err := api.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", currentUser)

supported features:
  - get user information
  - create, update, delete content
  - get comments, attachments, history, watchers  and children of content objects
  - get, add, delete labels
  - search using CQL

see https://github.com/coggsflod/confluence-go-api/tree/master/examples for more information and usage examples
*/
package goconfluence
