package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func makeHTML(body string) string {
	return "<html><body>" + body + "</body</html>"
}

func TestForEachNode(t *testing.T) {
	ts := []string{
		"",
		makeHTML(""),
		makeHTML("<p>word1 word2</p><script>\n        (function() {alert(1);})();\n    </script>"),
		makeHTML("<strong>this is test sentence.</strong><p>word1 word2\n<!-- this is comment -->\n</p>"),
		makeHTML(`<p>this is test sentence.</p><img src="https://example.com"/>`),
	}

	b := bytes.NewBuffer(nil)
	w = b
	for _, tc := range ts {
		t.Run("Case", func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(tc))
			if err != nil {
				t.Fatal(err)
			}

			forEachNode(doc, startElement, endElement)

			if _, err := html.Parse(b); err != nil {
				t.Errorf("unexpected result. err: %v", err)
			}
		})
	}
}
