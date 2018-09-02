package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	extractTextNode(doc)
}

func extractTextNode(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	if !(n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style")) {
		extractTextNode(n.FirstChild)
	}
	extractTextNode(n.NextSibling)
}
