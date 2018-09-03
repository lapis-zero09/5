package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

var (
	depth int
	w     io.Writer = os.Stdout
)

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
		if n.FirstChild == nil {
			fmt.Fprintf(w, " />\n")
		} else {
			fmt.Fprintf(w, ">\n")
			depth++
		}
	case html.TextNode:
		s := strings.TrimSpace(n.Data)
		if len(s) > 0 {
			fmt.Fprintf(w, "%*s%s\n", depth*2, "", s)
		}
	case html.CommentNode:
		fmt.Fprintf(w, "%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
