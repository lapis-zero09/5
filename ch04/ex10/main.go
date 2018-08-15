package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lapis-zero09/5/ch04/ex10/github"
)

func printIssues(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now().UTC()
	tOneMonthAgo := t.AddDate(0, -1, 0)
	tOneYearAgo := t.AddDate(-1, 0, 0)

	var oneMonthAgo, oneYearAgo, moreThanOneYearAgo []*github.Issue

	for _, item := range result.Items {
		if tOneMonthAgo.Before(item.CreatedAt) {
			oneMonthAgo = append(oneMonthAgo, item)
			continue
		}
		if tOneYearAgo.Before(item.CreatedAt) {
			oneYearAgo = append(oneYearAgo, item)
			continue
		}
		moreThanOneYearAgo = append(moreThanOneYearAgo, item)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	fmt.Println("-----------------一ヶ月以内---------------------")
	printIssues(oneMonthAgo)

	fmt.Println("-----------------一年以内---------------------")
	printIssues(oneYearAgo)

	fmt.Println("-----------------それ以上---------------------")
	printIssues(moreThanOneYearAgo)

}
