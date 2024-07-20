package parser

import (
	"golang.org/x/net/html"
)

func FetchCSRF(n *html.Node, csrf *string) {
	if *csrf != "" {
		return
	}

	if n.Type == html.ElementNode || n.Type == html.TextNode {
		for _, attr := range n.Attr {
			if attr.Key == "value" {
				*csrf = attr.Val
			}

		}

	}

	// Traverse the child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		FetchCSRF(c, csrf)
	}

}
