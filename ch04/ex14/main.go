package main

import (
	"flag"
	"log"
	"os"

	"html/template"

	"github.com/lapis-zero09/5/ch04/ex11/github"
)

var (
	issueTmpl = template.Must(template.ParseFiles("issueTmpl.html"))
	owner     = flag.String("o", "lapis-zero09", "repository owner")
	repo      = flag.String("r", "5", "repository name")
	issueId   = flag.Int("i", 0, "issue ID")
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal(os.Stderr, "set github token  like export GITHUB_TOKEN=\"xxxxxxxxxxxx\"\n")
		os.Exit(1)
	}
	reqUser := github.NewReqUser(token)

	flag.Parse()
	if *issueId == 0 {
		log.Fatal(os.Stderr, "set issue ID\n")
		os.Exit(1)
	}

	issue, err := reqUser.GetIssue(*owner, *repo, *issueId)
	if err != nil {
		log.Fatalf("error: %v\n", err)
		os.Exit(1)
	}

	if err := issueTmpl.Execute(os.Stdout, issue); err != nil {
		log.Fatal(err)
	}
}
