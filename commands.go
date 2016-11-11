package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandGet,
}

var commandGet cli.Command = cli.Command{
	Name:    "get",
	Aliases: []string{"g"},
	Usage:   "get backlog task",
	Action: func(c *cli.Context) error {
		issue_key := getIssueKeyFromString(c.Args().First())
		clone_from := c.Args().First()
		clone_to := os.Getenv("HOME") + "/.bhq/" + issue_key

		fmt.Println("    ", clone_from, "->", clone_to)
		_, err := exec.Command("mkdir", "-p", clone_to).Output()

		if err != nil {
			fmt.Println(err)
		}
		return nil
	},
}

// Convert url to issue key
func getIssueKeyFromString(url_str string) string {
	u, _ := url.Parse(url_str)
	issue_key := strings.Split(u.Path, "/")
	return issue_key[2]
}
