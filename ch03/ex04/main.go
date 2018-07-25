package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lapis-zero09/5/ch03/ex04/saddle"
)

func main() {
	// http://localhost:8000/?color=blue&height=1000&width=1000
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	saddleParam := map[string]string{
		"color":  "red",
		"height": "1000",
		"width":  "1000",
	}

	for key, val := range r.URL.Query() {
		if hasKeyInList(key, saddleParam) {
			fmt.Println(key, val[0])
			saddleParam[key] = val[0]
		}
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	saddle.Saddle(w, saddleParam)
}

func hasKeyInList(word string, dict map[string]string) bool {
	for key := range dict {
		if key == word {
			return true
		}
	}
	return false
}
