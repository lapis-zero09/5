package main

import (
	"fmt"
	"log"

	"github.com/lapis-zero09/5/ch04/ex11/github"
)

func printIssues(issues []*github.Issue) {
	for _, issue := range issues {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			issue.Number, issue.User.Login, issue.Title, issue.CreatedAt)
	}
}

func main() {
	// result, err := github.SearchIssues(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%d issues:\n", result.TotalCount)
	// printIssues(result.Items)

	// fmt.Println("-------------------------------------------")
	//
	// issues, err := github.GetIssues("lapis-zero09", "BrownBoost")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// printIssues(issues)

	// issue, err := github.GetIssue("lapis-zero09", "BrownBoost", 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("#%-5d %9.9s %.55s %v\n",
	// 	issue.Number, issue.User.Login, issue.Title, issue.CreatedAt)

	// issue, err := github.CreateIssue("lapis-zero09", "5", "issue from golang", "Hello World!\nthis issue made from lapis-zero09/5/ex11.")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("#%-5d %9.9s %.55s %v\n",
	// 	issue.Number, issue.User.Login, issue.Title, issue.CreatedAt)

	// issue, err := github.EditIssue("lapis-zero09", "5", 1, "edited from golang", "Hello World!\nthis issue made from lapis-zero09/5/ex11.")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("#%-5d %9.9s %.55s %v\n",
	// 	issue.Number, issue.User.Login, issue.Title, issue.CreatedAt)

	issue, err := github.CloseIssue("lapis-zero09", "5", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("#%-5d %9.9s %.55s %v\n",
		issue.Number, issue.User.Login, issue.Title, issue.CreatedAt)
}
