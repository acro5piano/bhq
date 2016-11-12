package main

import (
	//	"json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
//	"os/exec"
	"strings"

	"github.com/motemen/ghq/utils"
)

const BACKLOG_ROOT_URL = "https://scat919.backlog.jp/api/v2"

// Convert url to issue key
func GetIssueKeyFromURL(url_str string) string {
	u, err := url.Parse(url_str)
	utils.DieIf(err)

	issue_key := strings.Split(u.Path, "/")
	return issue_key[2]
}

func Whoami() string {
	return curl("users/myself")
}

func api_key() string {
	api_key := os.Getenv("BACKLOG_API_KEY")
	if api_key == "" {
		fmt.Println("please set BACKLOG_API_KEY")
		os.Exit(1)
	}

	return api_key
}

func curl(url string) string{
	resp, err := http.Get(fmt.Sprintf(
		"%s/%s?apiKey=%s",
		BACKLOG_ROOT_URL, url, api_key()))

	utils.DieIf(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func current_issue() string{
	pwd, _ := os.Getwd()
	return strings.Split(pwd, "/")[4]
}
