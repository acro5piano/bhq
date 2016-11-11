package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var Version string = "0.0.1"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "bhq"
	app.Usage = "Manage Backlog"
	app.Version = Version
	app.Author = "acro5piano"
	app.Email = "ketsume0211@gmail.com"
	app.Commands = Commands
	return app
}
