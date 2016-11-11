package main

import (
	"fmt"
	"os"
	"os/exec"

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
		issue_key := GetIssueKeyFromURL(c.Args().First())
		clone_from := c.Args().First()
		clone_to := os.Getenv("HOME") + "/.bhq/" + issue_key

		fmt.Println("    ", clone_from, "->", clone_to)
		_, err := exec.Command("mkdir", "-p", clone_to).Output()

		fmt.Println(GetTitle(issue_key))

		if err != nil {
			fmt.Println(err)
		}
		return nil
	},
}
