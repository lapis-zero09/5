package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func makeHTML(body string) string {
	return "<html><body>" + body + "</body</html>"
}

func TestCountWordsAndImages(t *testing.T) {
	ts := []struct {
		HTML     string
		expected [2]int
	}{
		{
			HTML:     makeHTML(""),
			expected: [2]int{0, 0},
		},
		{
			HTML:     makeHTML("<p>word1 word2</p>"),
			expected: [2]int{2, 0},
		},
		{
			HTML:     makeHTML("<strong>this is test sentence.</strong><p>word1 word2</p>"),
			expected: [2]int{6, 0},
		},
		{
			HTML:     makeHTML(`<p>this is test sentence.</p><img src="https://example.com" />`),
			expected: [2]int{4, 1},
		},
	}

	for _, tc := range ts {
		t.Run("Case", func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) { io.WriteString(w, tc.HTML) })
			s := httptest.NewServer(mux)
			defer s.Close()

			words, images, err := CountWordsAndImages(s.URL)
			if err != nil {
				t.Fatal(err)
			}
			if [2]int{words, images} != tc.expected {
				t.Errorf("unexpected result. expected: %v, but got: (%v,%v)", tc.expected, words, images)
			}
		})
	}
}
