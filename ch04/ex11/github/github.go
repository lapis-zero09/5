package github

import "time"

const endPoint = "https://api.github.com"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
	Milestone *Milestone
	Assignees []*User
}

type Milestone struct {
	Title string    `json:"title"`
	DueOn time.Time `json:"due_on"`
}

type User struct {
	Login     string
	HTMLURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
}
