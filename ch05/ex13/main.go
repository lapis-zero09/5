package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(host, u string) []string {
	fmt.Println(u)
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		err = fmt.Errorf("getting %s: %s", u, resp.Status)
		log.Fatal(err)
		return nil
	}

	parseURL, err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
	}
	if parseURL.Host == host {
		fpath := makePath(parseURL)
		os.MkdirAll(filepath.Dir(fpath), os.ModePerm)

		f, err := os.Create(fpath)
		if err == nil {
			defer f.Close()
		} else {
			log.Print(err)
		}
	}

	list, err := extract(resp)
	if err != nil {
		log.Print(err)
	}
	return list
}

func makePath(u *url.URL) string {
	p := u.Path
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	if strings.HasSuffix(p, "/") {
		p = p + "index.html"
	}
	return filepath.FromSlash("./" + u.Host + p)
}

func extract(resp *http.Response) ([]string, error) {
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
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

func main() {
	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	breadthFirst(func(item string) []string {
		return crawl(u.Host, item)
	}, os.Args[1:])
}
