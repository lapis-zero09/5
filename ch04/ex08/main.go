package main

import (
	"fmt"
	"unicode"
)

func printCategoryCounts(s string) {
	categories := map[string]func(rune) bool{
		// ref: https://www.compart.com/en/unicode/category
		"C": unicode.IsControl,
		"L": unicode.IsLetter,
		"M": unicode.IsMark,
		"N": unicode.IsNumber,
		"P": unicode.IsPunct,
		"S": unicode.IsSymbol,
		"Z": unicode.IsSpace,
	}
	cntsCategories := make(map[string]int)
	cnts := make(map[rune]int)
	for _, r := range s {
		cnts[r]++
	}

	fmt.Println("term\tcount")
	for k, v := range cnts {
		fmt.Printf("%#U \t %d\n", k, v)
		for category, f := range categories {
			if f(k) {
				cntsCategories[category] += v
			}
		}
	}

	fmt.Println("category\tcount")
	for category, count := range cntsCategories {
		fmt.Printf("%s \t %d\n", category, count)
	}
}

func main() {
	printCategoryCounts("aaaああ　　  \t亜\n")
}
