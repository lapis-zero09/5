package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

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

	id := "playground"
	n := ElementByID(doc, id)
	if n != nil {
		s := fmt.Sprintf("<%s", n.Data)
		for _, a := range n.Attr {
			s += fmt.Sprintf(" %s='%s'", a.Key, a.Val)
		}
		s += ">\n"
		fmt.Println(s)
	}
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil {
		if !post(n) {
			return false
		}
	}

	return true
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node
	forEachNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key != "id" {
					continue
				}
				if a.Val != id {
					continue
				}
				node = n
				return false
			}
		}
		return true
	}, nil)

	return node
}
