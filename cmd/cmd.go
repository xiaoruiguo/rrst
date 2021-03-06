package cmd

import (
	"fmt"
	"github.com/catay/rrst/app"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

// the version string will be injected during the build.
var Version = "0.0.0"

const (
	author        = "Steven Mertens <steven.mertens@catay.be>"
	defaultConfig = "/etc/rsst.yaml"
)

type cli struct {
	app          *kingpin.Application
	configFile   *string
	verbose      *bool
	listCmd      *kingpin.CmdClause
	syncCmd      *kingpin.CmdClause
	showCmd      *kingpin.CmdClause
	cleanCmd     *kingpin.CmdClause
	syncArgRepo  *string
	cleanArgRepo *string
}

func New() *cli {
	c := new(cli)
	c.app = kingpin.New("rsst", "Remote Repository Sync Tool")
	c.app.Version(Version)
	c.app.Author(author)
	c.configFile = c.app.Flag("config", "Path to alternate YAML configuration file.").Short('c').Default(defaultConfig).String()
	c.verbose = c.app.Flag("verbose", "Turn on verbose output. Default is verbose turned off.").Short('v').Bool()
	c.listCmd = c.app.Command("list", "List repository names and description.")
	c.syncCmd = c.app.Command("sync", "Synchronize remote to local repository sets.")
	c.syncArgRepo = c.syncCmd.Arg("repo name", "Repository name.").String()
	c.showCmd = c.app.Command("show", "Show available repository sets.")
	c.cleanCmd = c.app.Command("clean", "Cleanup repository cache.")
	c.cleanArgRepo = c.cleanCmd.Arg("repo name", "Repository name.").String()

	return c
}

func (self *cli) Run() (err error) {
	args := kingpin.MustParse(self.app.Parse(os.Args[1:]))
	r, err := app.NewApp(*self.configFile)
	if err != nil {
		return err
	}

	switch args {
	case "sync":
		if err := r.Sync(*self.syncArgRepo); err != nil {
			return err
		}
	case "show":
		fmt.Println("command: show")
		r.Show()
	case "list":
		r.List()
	case "clean":
		if err := r.Clean(*self.cleanArgRepo); err != nil {
			return err
		}
	}

	return nil
}
