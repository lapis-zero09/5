package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lapis-zero09/5/ch03/ex09/fractal"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fractalParams := map[string]int{
		"x":             2,
		"y":             2,
		"magnification": 1024,
	}

	for key, val := range r.URL.Query() {
		parseValue, err := strconv.Atoi(val[0])
		if err != nil {
			continue
		}
		if hasKeyInList(key, fractalParams) {
			fmt.Println(key, parseValue)
			fractalParams[key] = parseValue
		}
	}

	fractal.Fractal(w, fractalParams)
}

func hasKeyInList(word string, dict map[string]int) bool {
	for key := range dict {
		if key == word {
			return true
		}
	}
	return false
}
