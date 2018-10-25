package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Selector interface {
	Match(xml.StartElement) bool
	GetVal() string
}

type Selectors interface {
	GetSelectors() string
}

type TypeSelector struct {
	Value string
}

type AttrSelector struct {
	Name  string
	Value string
}

type SelectorList []Selector

func (ts *TypeSelector) Match(elm xml.StartElement) bool {
	return ts.Value == elm.Name.Local
}

func (as *AttrSelector) Match(elm xml.StartElement) bool {
	val, ok := getAttr(elm, as.Name)
	return val == as.Value && ok
}

func (ts *TypeSelector) GetVal() string {
	return ts.Value
}

func (as *AttrSelector) GetVal() string {
	return as.Value
}

func (selectors SelectorList) GetSelectors() string {
	s := ""
	for idx, selector := range selectors {
		if idx != 0 {
			s += " "
		}
		s += selector.GetVal()
	}
	return s
}

func getAttr(elm xml.StartElement, name string) (string, bool) {
	for _, attr := range elm.Attr {
		if attr.Name.Local != name {
			continue
		}
		return attr.Value, true
	}
	return "", false
}

func parse(s string) (Selector, error) {
	switch {
	case strings.HasPrefix(s, "."):
		return &AttrSelector{"id", s[1:]}, nil

	case strings.HasPrefix(s, "#"):
		return &AttrSelector{"class", s[1:]}, nil

	default:
		return &TypeSelector{s}, nil
	}
}

func argParse(args []string) (SelectorList, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("selector is empty!")
	}

	selectors := make(SelectorList, 0)
	for _, arg := range args {
		selector, err := parse(arg)
		if err != nil {
			return nil, err
		}
		selectors = append(selectors, selector)
	}
	return selectors, nil
}

func containsAll(elems []xml.StartElement, selectors SelectorList) bool {
	for len(selectors) <= len(elems) {
		if len(selectors) == 0 {
			return true
		}
		if selectors[0].Match(elems[0]) {
			selectors = selectors[1:]
		}
		elems = elems[1:]
	}
	return false
}

func main() {
	selectors, err := argParse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, selectors) {
				fmt.Printf("%s: %s\n", selectors.GetSelectors(), tok)
			}
		}
	}
}
