package main

import (
	"net/url"
	"os"
	"strings"
)

// Convert url to issue key
func GetIssueKeyFromURL(url_str string) string {
	u, _ := url.Parse(url_str)
	issue_key := strings.Split(u.Path, "/")
	return issue_key[2]
}

func GetTitle(issue_key string) string {
	return api_key()
}

func api_key() string {
	return os.Getenv("BACKLOG_API_KEY")
}
