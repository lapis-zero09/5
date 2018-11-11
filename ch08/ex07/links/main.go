package links

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

func Extract(u string) ([]string, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", u, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", u, err)
	}

	filePath := path.Base(resp.Request.URL.Path)
	f, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("file creating error:%v", err)
	}
	defer f.Close()

	// htmlではない場合探索しない
	if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		if _, err = io.Copy(f, resp.Body); err != nil {
			return nil, fmt.Errorf("file writing error: %v", err)
		}
		return nil, nil
	}

	html.Render(f, doc)

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

				if link.Host != resp.Request.URL.Host {
					continue
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
