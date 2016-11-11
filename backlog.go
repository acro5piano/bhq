package main

import (
	"net/url"
	"strings"
)

// Convert url to issue key
func GetIssueKeyFromURL(url_str string) string {
	u, _ := url.Parse(url_str)
	issue_key := strings.Split(u.Path, "/")
	return issue_key[2]
}

func GetTitle(issue_key string) string {
	return "title"
}
