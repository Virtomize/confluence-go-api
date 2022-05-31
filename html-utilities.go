package goconfluence

import (
	"bytes"
	"errors"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func getBody(doc *html.Node) (*html.Node, error) {
	var b *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && strings.ToLower(n.Data) == "body" {
			b = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if b != nil {
		return b, nil
	}
	return nil, errors.New("missing <body> in the node tree")
}

func stripImgs(doc *html.Node) *html.Node {
	var f func(*html.Node, *html.Node)
	f = func(n, parent *html.Node) {
		if parent != nil && n.Type == html.ElementNode && (strings.ToLower(n.Data) == "img" || strings.ToLower(n.Data) == "script") {
			parent.RemoveChild(n)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c, n)
			}
		}
	}
	f(doc, nil)
	return doc
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	err := html.Render(w, n)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

// StripHTML removes the specified information from the HTML and sets it as a string
func StripHTML(buf []byte, bodyOnly, stripImg bool) string {
	doc, err := html.Parse(bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}
	if bodyOnly {
		doc, err = getBody(doc)
		if err != nil {
			log.Fatal(err)
		}
	}
	if stripImg {
		doc = stripImgs(doc)
	}
	result := renderNode(doc)
	return result
}
