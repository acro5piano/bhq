package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/motemen/ghq/utils"
	"github.com/antonholmquist/jason"

)

const BACKLOG_ROOT_URL = "https://scat919.backlog.jp/api/v2"
const TMP_FILE = "/tmp/backlog-comment"

type Issue struct {
	Key string
	Summary string
}


// Convert url to issue key
func GetIssueKeyFromURL(url_str string) string {
	u, err := url.Parse(url_str)
	utils.DieIf(err)

	issue_key := strings.Split(u.Path, "/")
	return issue_key[2]
}

func GetSummary(key string) string{
	v, err := jason.NewObjectFromBytes(backlog_get_byte("issues/"+key))
	utils.DieIf(err)
	summary, err := v.GetString("summary")
	utils.DieIf(err)
	return summary
}

func Whoami() string {
	return backlog_get("users/myself")
}

func AddComment() string {
	edit_with_external_editor(TMP_FILE)
	data, err := ioutil.ReadFile(TMP_FILE)
	utils.DieIf(err)
	os.Remove(TMP_FILE)

	response := backlog_post("issues/"+current_issue()+"/comments",
		url.Values{"content": {string(data)}})
	return response
}

func ListIssues() []Issue {
	issues := []Issue{}

	files, _ := ioutil.ReadDir(bhq_root())
	for _, f := range files {
		issues = append(issues, Issue{
			Key: f.Name(),
			Summary: GetSummary(f.Name()),
		})
	}
	return issues
}


func api_key() string {
	api_key := os.Getenv("BACKLOG_API_KEY")
	if api_key == "" {
		fmt.Println("please set BACKLOG_API_KEY")
		os.Exit(1)
	}

	return api_key
}

func backlog_get(url string) string{
	resp, err := http.Get(fmt.Sprintf(
		"%s/%s?apiKey=%s",
		BACKLOG_ROOT_URL, url, api_key()))

	utils.DieIf(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func backlog_get_byte(url string) []byte{
	resp, err := http.Get(fmt.Sprintf(
		"%s/%s?apiKey=%s",
		BACKLOG_ROOT_URL, url, api_key()))

	utils.DieIf(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func backlog_post(url string, values url.Values) string{
	resp, err := http.PostForm(fmt.Sprintf(
		"%s/%s?apiKey=%s", BACKLOG_ROOT_URL, url, api_key()),
		values)

	utils.DieIf(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func current_issue() string{
	pwd, _ := os.Getwd()
	return strings.Split(pwd, "/")[4]
}

func edit_with_external_editor(file_path string) error{
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = os.Getenv("VISUAL")
	}
	cmd := exec.Command(editor, file_path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return nil
}
