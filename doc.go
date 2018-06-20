/*
Package goconfluence implementing atlassian's Confluence API

Initialize a new API instance

	api, err := goconfluence.NewAPI("https://<your-domain>.atlassian.net", "<username>", "<api-token>")
	if err != nil {
		log.Fatal(err)
	}

supported features:
	- get user information
	- create, update, delete content
	- get comments, attachments and children of content objects
	- search using CQL

see https://github.com/cseeger-epages/confluence-go-api/tree/master/examples for more information and usage examples

*/
package goconfluence
