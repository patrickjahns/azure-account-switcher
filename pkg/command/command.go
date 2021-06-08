package command

import (
	"os"

	"github.com/patrickjahns/azure-account-switcher/pkg/version"
	"github.com/urfave/cli/v2"
)

// Run parses the command line arguments and executes the program.
func Run() error {

	app := &cli.App{
		Name:    "aas",
		Version: version.Info(),
		Usage:   "Azure Account Switcher",
		Authors: []*cli.Author{
			{
				Name:  "Patrick Jahns",
				Email: "github@patrickjahns.de",
			},
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show help",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Prints the current version",
	}

	app.Action = func(c *cli.Context) error {
		return run()
	}

	return app.Run(os.Args)
}

func run() error {
	return nil
}
