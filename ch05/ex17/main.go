package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	doc, err := getPage(url)
	if err != nil {
		log.Fatal(err)
	}

	nodes, err := ElementsByTagName(doc, os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range nodes {
		fmt.Println(n.Data)
	}
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

func ElementsByTagName(doc *html.Node, queries ...string) ([]*html.Node, error) {
	var nodeArray []*html.Node
	forEachNode(doc,
		func(n *html.Node) {
			if n.Type == html.ElementNode {
				if include(n.Data, queries) {
					nodeArray = append(nodeArray, n)
				}
			}
		},
		nil)
	return nodeArray, nil
}

func include(v string, vals []string) bool {
	for _, val := range vals {
		if v == val {
			return true
		}
	}
	return false
}

func getPage(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
