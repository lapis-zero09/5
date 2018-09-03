package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func makeHTML(body string) string {
	return "<html><body>" + body + "</body</html>"
}

func TestElementByID(t *testing.T) {
	ts := []struct {
		HTML     string
		ID       string
		expected string
	}{
		{
			HTML:     makeHTML(""),
			ID:       "",
			expected: "",
		},
		{
			HTML:     makeHTML(""),
			ID:       "aaa",
			expected: "",
		},

		{
			HTML:     makeHTML("<p id='test' class='aaa'>word1 word2</p>"),
			ID:       "test",
			expected: "p id='test' class='aaa'",
		},
		{
			HTML:     makeHTML("<p id='test' class='aaa'>word1 word2</p>"),
			ID:       "not test",
			expected: "",
		},
		{
			HTML:     makeHTML("<strong>this is test sentence.</strong><p>word1 word2</p><img id='a' src='./cc'/>"),
			ID:       "a",
			expected: "img id='a' src='./cc'",
		},
	}

	for _, tc := range ts {
		t.Run("Case", func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(tc.HTML))
			if err != nil {
				t.Fatal(err)
			}

			n := ElementByID(doc, tc.ID)
			s := ""
			if n != nil {
				s += n.Data
				for _, a := range n.Attr {
					s += fmt.Sprintf(" %s='%s'", a.Key, a.Val)
				}
			}

			if s != tc.expected {
				t.Errorf("unexpected result. expected: (%s), but got: (%s)", tc.expected, s)
			}
		})
	}
}
