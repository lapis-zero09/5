package github

import (
	"bytes"
	"encoding/json"
	"path"
	"strconv"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func (r *ReqUser) SearchIssues(terms []string) (*IssuesSearchResult, error) {
	req, err := r.newRequest("GET", "/search/issues", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("q", strings.Join(terms, " "))
	req.URL.RawQuery = q.Encode()

	var result *IssuesSearchResult
	if err := r.do(req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ReqUser) GetIssues(owner, repo string) ([]*Issue, error) {
	// ref: https://developer.github.com/v3/issues/#list-issues-for-a-repository
	uri := path.Join("/repos", owner, repo, "issues")
	req, err := r.newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	var issues []*Issue
	if err := r.do(req, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func (r *ReqUser) GetIssue(owner, repo string, number int) (*Issue, error) {
	// ref: https://developer.github.com/v3/issues/#get-a-single-issue
	uri := path.Join("/repos", owner, repo, "issues", strconv.Itoa(number))
	req, err := r.newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	var issue *Issue
	if err := r.do(req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}

func (r *ReqUser) CreateIssue(owner, repo, subject, content string) (*Issue, error) {
	// TOr.DO: milestone, labels, assignees
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
	req, err := r.newRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	var issue *Issue
	if err := r.do(req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}

func (r *ReqUser) EditIssue(owner, repo string, number int, subject, content string) (*Issue, error) {
	// TOr.DO: milestone, labels, assignees
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
	req, err := r.newRequest("PATCH", uri, body)
	if err != nil {
		return nil, err
	}

	var issue *Issue
	if err := r.do(req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}

func (r *ReqUser) CloseIssue(owner, repo string, number int) (*Issue, error) {
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
	req, err := r.newRequest("PATCH", uri, body)
	if err != nil {
		return nil, err
	}

	var issue *Issue
	if err := r.do(req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}
