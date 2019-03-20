# confluence-go-api

[![GoDoc](https://img.shields.io/badge/godoc-reference-green.svg)](https://godoc.org/github.com/cseeger-epages/confluence-go-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/cseeger-epages/confluence-go-api)](https://goreportcard.com/report/github.com/cseeger-epages/confluence-go-api)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/cseeger-epages/confluence-go-api/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/cseeger-epages/confluence-go-api.svg?branch=master)](https://travis-ci.org/cseeger-epages/confluence-go-api)
[![Built with Mage](https://magefile.org/badge.svg)](https://magefile.org)


is a [Confluence](https://www.atlassian.com/software/confluence) REST API client implementation written in [GOLANG](https://golang.org).

## Supportet Features

- get, update, delete content
- get comments, attachments, children of content objects, history, watchers
- get, add ,delete labels
- get user information
- search using [CQL](https://developer.atlassian.com/cloud/confluence/advanced-searching-using-cql/)

If you miss some feature implementation, feel free to open an issue or send pull requests. I will take look as soon as possible.

## Installation

If you already installed GO on your system and configured it properly than its simply:

```
go get github.com/cseeger-epages/confluence-go-api
```

If not follow [these instructions](https://nats.io/documentation/tutorials/go-install/).

## Usage

### Simple example

```
package main

import (
  "fmt"
  "log"

  "github.com/cseeger-epages/confluence-go-api"
)

func main() {

  // initialize a new api instance
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
}
```

### Advanced examples

see [examples](https://github.com/cseeger-epages/confluence-go-api/tree/master/examples) for some more usage examples

## Code Documentation

You find the full [code documentation here](https://godoc.org/github.com/cseeger-epages/confluence-go-api).

The Confluence API documentation [can be found here](https://docs.atlassian.com/ConfluenceServer/rest/6.9.1/).
