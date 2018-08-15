package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/lapis-zero09/5/ch04/ex11/github"
)

func printIssue(issue *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s %v\n",
		issue.Number, issue.User.Login, issue.Title, issue.CreatedAt)
}

func printIssues(issues []*github.Issue) {
	for _, issue := range issues {
		printIssue(issue)
	}
}

func useEditor(editor string) ([][]byte, error) {
	tmpFile, err := ioutil.TempFile("", "5_ch04_ex11_")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	cmd := exec.Command(editor, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		return nil, err
	}

	return bytes.SplitN(b, []byte("\n"), 2), nil
}

var (
	owner   = flag.String("o", "lapis-zero09", "repository owner")
	repo    = flag.String("r", "5", "repository name")
	editor  = flag.String("e", "vim", "use editor")
	issueId = flag.Int("i", 0, "issue ID\nmehtod:edit/close/(get)")
	method  = flag.String("m", "", "Chose method for issue - get/create/edit/close")
	subject = flag.String("s", "", "Subject for issue\nmehtod:create/edit")
	content = flag.String("c", "", "Contet for issue\nmehtod:create/edit")
)

func main() {
	flag.Parse()
	if *method == "" {
		fmt.Fprintf(os.Stderr, "Should Chose method for issue - get/create/edit/close\n")
		os.Exit(1)
	}
	fmt.Printf("acess to https://github.com/%s/%s", *owner, *repo)

	switch *method {
	case "get":
		if *issueId == 0 {
			issues, err := github.GetIssues(*owner, *repo)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Your request is successfully done.")
			printIssues(issues)
		} else {
			issue, err := github.GetIssue(*owner, *repo, *issueId)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Your request is successfully done.")
			printIssue(issue)
		}

	case "create":
		if *subject == "" {
			bs, err := useEditor(*editor)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				fmt.Println("use -s subject and -c content / -e editor")
				os.Exit(1)
			}
			*subject, *content = string(bs[0]), string(bs[1])
		}

		issue, err := github.CreateIssue(*owner, *repo, *subject, *content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Your request is successfully done.")
		printIssue(issue)

	case "edit":
		if *issueId == 0 {
			fmt.Fprintf(os.Stderr, "Should Enter issueID -i param\n")
			os.Exit(1)
		}

		if *subject == "" && *content == "" {
			bs, err := useEditor(*editor)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				fmt.Println("use -s subject and -c content / -e editor")
				os.Exit(1)
			}
			*subject, *content = string(bs[0]), string(bs[1])
		}

		issue, err := github.EditIssue(*owner, *repo, *issueId, *subject, *content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Your request is successfully done.")
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			issue.Number, issue.User.Login, issue.Title, issue.CreatedAt)

	case "close":
		if *issueId == 0 {
			fmt.Fprintf(os.Stderr, "Should Enter issueID -i param\n")
			os.Exit(1)
		}

		issue, err := github.CloseIssue(*owner, *repo, *issueId)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Your request is successfully done.")
		printIssue(issue)
	}

	// 実行したコマンドの結果を出力
	// fmt.Printf("ls result: \n%s", string(out))
	// result, err := github.SearchIssues(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%d issues:\n", result.TotalCount)
	// printIssues(result.Items)

}
