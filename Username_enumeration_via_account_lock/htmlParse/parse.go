package htmlparse

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func HtmlTraversal(n *html.Node, username string, foundUser *string) {

	if n.Type == html.ElementNode || n.Type == html.TextNode {
		if strings.Contains(n.Data, "You have made too many incorrect login attempts.") {
			//fmt.Println("---------------->FOUND FOUND FOUND FOUND FOUND FOUND FOUND FOUND<---------------- ", username)
			*foundUser = username
			return
		}
	}
	// Traverse the child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		HtmlTraversal(c, username, foundUser)
	}
	return
}

func HtmlTraversalPassword(n *html.Node, password *string, foundPassword *string) {

	// if n.Type == html.ElementNode || n.Type == html.TextNode {
	// 	// if strings.Contains(n.Data, "You have made too many incorrect login attempts. Please try again in 1 minute(s).") {
	// 	// 	fmt.Println("Incorrect Password")
	// 	// 	*foundPassword = *password
	// 	// } else {
	// 	// 	fmt.Println("Correct Password")
	// 	// 	*foundPassword = *password
	// 	// 	return
	// 	// }
	// 	fmt.Println(n.Data)
	// }

	if n.Type == html.ElementNode || n.Type == html.TextNode {
		if n.Type == html.ElementNode {
			*foundPassword += fmt.Sprintf("<%s", n.Data)
			for _, attr := range n.Attr {
				// if !strings.Contains(attr.Val, "is-warning") {
				// 	fmt.Println("Password Found")
				// } else {
				// 	fmt.Println("I have a Warning")
				// }
				*foundPassword += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
			}
			*foundPassword += ">\n"
		}
		// if n.Type == html.TextNode {
		// 	fmt.Print(n.Data)
		// }
		// if n.Type == html.ElementNode {
		// 	fmt.Printf("</%s>", n.Data)
		// }
	}

	// Traverse the child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		HtmlTraversalPassword(c, password, foundPassword)
	}

	return
}
