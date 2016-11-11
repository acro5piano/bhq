package main

import (
	//	"json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/motemen/ghq/utils"
)

const BACKLOG_ROOT_URL = "https://scat919.backlog.jp/api/v2/users/myself?apiKey="

// Convert url to issue key
func GetIssueKeyFromURL(url_str string) string {
	u, _ := url.Parse(url_str)
	issue_key := strings.Split(u.Path, "/")
	return issue_key[2]
}

func GetTitle(issue_key string) string {
	resp, err := http.Get("https://scat919.backlog.jp/api/v2/users/myself?apiKey=" + api_key())
	utils.DieIf(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func api_key() string {
	api_key := os.Getenv("BACKLOG_API_KEY")
	if api_key == "" {
		os.Exit(1)
	}

	return api_key
}
