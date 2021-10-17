package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMMLURL  string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // Markdown
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
