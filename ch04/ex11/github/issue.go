package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
)

var endPoint = "https://api.github.com"
var token = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

func newRequest(method, uri string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, endPoint+uri, body)
	if err != nil {
		return nil, err
	}
	// req.Header.Set(
	// 	"Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", "token "+token)
	return req, nil
}

func do(req *http.Request, i interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 2xx Success系ではないもの
	if resp.StatusCode > 299 {
		return fmt.Errorf("failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(i); err != nil {
		return err
	}
	return nil
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	req, err := newRequest("GET", "/search/issues", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("q", strings.Join(terms, " "))
	req.URL.RawQuery = q.Encode()

	var result *IssuesSearchResult
	if err := do(req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetIssues(owner, repo string) ([]*Issue, error) {
	// ref: https://developer.github.com/v3/issues/#list-issues-for-a-repository
	uri := path.Join("/repos", owner, repo, "issues")
	req, err := newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	var issues []*Issue
	if err := do(req, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func GetIssue(owner, repo string, number int) (*Issue, error) {
	// ref: https://developer.github.com/v3/issues/#get-a-single-issue
	uri := path.Join("/repos", owner, repo, "issues", strconv.Itoa(number))
	req, err := newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	var issue *Issue
	if err := do(req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}

func CreateIssue(owner, repo, subject, content string) (*Issue, error) {
	// TODO: milestone, labels, assignees
	// ref: https://developer.github.com/v3/issues/#create-an-issue
	newIssue := struct {
		Subject string `json:"title"`
		Content string `json:"body"`
	}{
		Subject: subject,
		Content: content,
	}
	newIssueByte, err := json.Marshal(newIssue)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(newIssueByte)

	uri := path.Join("/repos", owner, repo, "issues")
	req, err := newRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	var issue *Issue
	if err := do(req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}

func EditIssue(owner, repo string, number int, subject, content string) (*Issue, error) {
	// TODO: milestone, labels, assignees
	// ref: https://developer.github.com/v3/issues/#edit-an-issue
	editIssue := struct {
		Subject string `json:"title"`
		Content string `json:"body"`
	}{
		Subject: subject,
		Content: content,
	}
	editIssueByte, err := json.Marshal(editIssue)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(editIssueByte)

	uri := path.Join("/repos", owner, repo, "issues", strconv.Itoa(number))
	req, err := newRequest("PATCH", uri, body)
	if err != nil {
		return nil, err
	}

	var issue *Issue
	if err := do(req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}

func CloseIssue(owner, repo string, number int) (*Issue, error) {
	// ref: https://developer.github.com/v3/issues/#edit-an-issue
	c := struct {
		State string `json:"state"`
	}{
		State: "closed",
	}
	cByte, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(cByte)

	uri := path.Join("/repos", owner, repo, "issues", strconv.Itoa(number))
	req, err := newRequest("PATCH", uri, body)
	if err != nil {
		return nil, err
	}

	var issue *Issue
	if err := do(req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}
