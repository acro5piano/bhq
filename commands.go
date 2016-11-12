package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
	"github.com/motemen/ghq/utils"
)

var Commands = []cli.Command{
	commandGet,
	commandWhoami,
	commandList,
}

var commandGet cli.Command = cli.Command{
	Name:    "get",
	Aliases: []string{"g"},
	Usage:   "get backlog task",
	Action: func(c *cli.Context) error {
		issue_key := GetIssueKeyFromURL(c.Args().First())
		clone_from := c.Args().First()
		clone_to := bhq_root() + "/" +  issue_key

		fmt.Println("    ", clone_from, "->", clone_to)
		_, err := exec.Command("mkdir", "-p", clone_to).Output()

//		fmt.Println(GetTitle(issue_key))

		if err != nil {
			fmt.Println(err)
		}
		return nil
	},
}
var commandWhoami cli.Command = cli.Command{
	Name:    "whoami",
	Aliases: []string{"w"},
	Usage:   "get backlog user",
	Action: func(c *cli.Context) error {
		fmt.Println(Whoami())
		return nil
	},
}
var commandList cli.Command = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "list backlog issues",
	Action: func(c *cli.Context) error {
		out, err := exec.Command("ls", bhq_root()).Output()
		utils.DieIf(err)
		fmt.Println(string(out))

		return nil
	},
}

func bhq_root() string{
	return os.Getenv("HOME") + "/.bhq"
}
