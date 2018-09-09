package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func makeHTML(body string) string {
	return "<html><body>" + body + "</body</html>"
}

func TestElementsByTagName(t *testing.T) {
	stringHTML := makeHTML("<div><h3>fdsaf</h3><p>brfds</p></div><h1>aaa</h1><h2>freafe</h2><p>fberaa</p><div><h3>fdsaf</h3><a href='http://example.com'></div><script>fdsa</script>")
	ts := []struct {
		html     string
		queries  []string
		expected string
	}{
		{
			html:     "",
			queries:  []string{},
			expected: "",
		},
		{
			html:     "",
			queries:  []string{"a"},
			expected: "",
		},
		{
			html:     stringHTML,
			queries:  []string{},
			expected: "",
		},
		{
			html:     stringHTML,
			queries:  []string{"a"},
			expected: "a",
		},
		{
			html:     stringHTML,
			queries:  []string{"h1", "h2", "h3"},
			expected: "h3h1h2h3",
		},
		{
			html:     stringHTML,
			queries:  []string{"h1", "p", "h3"},
			expected: "h3ph1ph3",
		},
		{
			html:     stringHTML,
			queries:  []string{"div", "p", "h3"},
			expected: "divh3ppdivh3",
		},
	}

	for _, tc := range ts {
		t.Run("Case", func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(tc.html))
			if err != nil {
				t.Error(err)
			}
			nodes, err := ElementsByTagName(doc, tc.queries...)

			s := ""
			for _, node := range nodes {
				s += node.Data
			}

			if err != nil {
				t.Error(err)
			}

			if s != tc.expected {
				t.Errorf("unexpected result. expected: (%s), but got: (%s)", tc.expected, s)
			}
		})
	}
}
