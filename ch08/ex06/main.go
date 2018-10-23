package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lapis-zero09/5/ch05/lesson06/links"
)

var tokens = make(chan struct{}, 20)

type link struct {
	url   string
	depth int
}

type linkList struct {
	list  []string
	depth int
}

func crawl(l link) linkList {
	fmt.Println(l.url)
	tokens <- struct{}{}
	list, err := links.Extract(l.url)
	<-tokens

	if err != nil {
		log.Print(err)
	}

	return linkList{list, l.depth}
}

//!-sema

//!+
func main() {
	depth := flag.Int("depth", 1, "depth")
	flag.Parse()

	var n int
	n++

	worklist := make(chan linkList)
	go func() { worklist <- linkList{flag.Args(), -1} }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, li := range list.list {
			if list.depth >= *depth {
				continue
			}
			if !seen[li] {
				seen[li] = true
				n++
				go func(l link) {
					worklist <- crawl(l)
				}(link{li, list.depth + 1})
			}
		}
	}
}
