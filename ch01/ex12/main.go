package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/lapis-zero09/5/ch01/ex12/lissajous"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajousParam := map[string]float64{
		"cycles":  5,
		"res":     0.001,
		"size":    100,
		"nframes": 64,
		"delay":   8,
	}

	for key, val := range r.URL.Query() {
		parseValue, err := strconv.ParseFloat(val[0], 64)
		if err != nil {
			continue
		}
		if hasKeyInList(key, lissajousParam) {
			fmt.Println(key, parseValue)
			lissajousParam[key] = parseValue
		}
	}
	lissajous.Lissajous(w, lissajousParam)
	// printLissajousParam(w, lissajousParam)
}

// func printLissajousParam(w http.ResponseWriter, lissajousParam map[string]float64) {
// 	for key, val := range lissajousParam {
// 		fmt.Fprintf(w, "%s = %f\n", key, val)
// 	}
// }

func hasKeyInList(word string, dict map[string]float64) bool {
	for key := range dict {
		if key == word {
			return true
		}
	}
	return false
}
