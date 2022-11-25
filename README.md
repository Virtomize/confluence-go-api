# confluence-go-api

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=VBXHBYFU44T5W&source=url)
[![GoDoc](https://img.shields.io/badge/godoc-reference-green.svg)](https://godoc.org/github.com/virtomize/confluence-go-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/virtomize/confluence-go-api)](https://goreportcard.com/report/github.com/virtomize/confluence-go-api)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/virtomize/confluence-go-api/blob/master/LICENSE)
[![Built with Mage](https://magefile.org/badge.svg)](https://magefile.org)


is a [Confluence](https://www.atlassian.com/software/confluence) REST API client implementation written in [GOLANG](https://golang.org).

## Supported Features

- get, update, delete content
- get, update, delete content templates and blueprints
- get comments, attachments, children of content objects, history, watchers
- get, add ,delete labels
- get user information
- search using [CQL](https://developer.atlassian.com/cloud/confluence/advanced-searching-using-cql/)

If you miss some feature implementation, feel free to open an issue or send pull requests. I will take look as soon as possible.

## Donation
If this project helps you, feel free to give us a cup of coffee :).

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=VBXHBYFU44T5W&source=url)

## Installation

If you already installed GO on your system and configured it properly than its simply:

```
go get github.com/virtomize/confluence-go-api
```

If not follow [these instructions](https://golang.org/doc/install)

## Usage

### Simple example

```
package main

import (
  "fmt"
  "log"

  "github.com/virtomize/confluence-go-api"
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


### Using a Personal Access Token

To generate a confluence personal access token (PAT) see this article: [using personal access tokens](https://confluence.atlassian.com/enterprise/using-personal-access-tokens-1026032365.html). Only set the token in the NewAPI function

```
  api, err := goconfluence.NewAPI("https://<your-domain>.atlassian.net/wiki/rest/api", "", "<personal-access-token>")
```

### Advanced examples

see [examples](https://github.com/virtomize/confluence-go-api/tree/master/examples) for some more usage examples

## Code Documentation

You find the full [code documentation here](https://godoc.org/github.com/virtomize/confluence-go-api).

The Confluence API documentation [can be found here](https://docs.atlassian.com/ConfluenceServer/rest/6.9.1/).

## Contribution

Thank you for participating to this project.
Please see our [Contribution Guidlines](https://github.com/virtomize/confluence-go-api/blob/master/CONTRIBUTING.md) for more information.

### Pre-Commit

This repo uses [pre-commit hooks](https://pre-commit.com/). Please install pre-commit and do `pre-commit install`

### Conventional Commits

Format commit messaged according to [Conventional Commits standard](https://www.conventionalcommits.org/en/v1.0.0/).

### Semantic Versioning

Whenever you need to version something make use of [Semantic Versioning](https://semver.org).


