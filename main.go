package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var buildCommit string // build number set at compile-time

func main() {
	logrus.Info(fmt.Printf("Drone Rubygems Plugin built from %s", buildCommit))

	app := cli.NewApp()
	app.Name = "rubygems drone plugin"
	app.Usage = "rubygems drone plugin"
	app.Action = run
	app.Version = buildCommit
	app.Flags = []cli.Flag{

		//
		// config args
		//

		cli.StringFlag{
			Name:   "api_key",
			Usage:  "rubygems api key",
			EnvVar: "RUBYGEMS_API_KEY,PLUGIN_API_KEY",
		},
		cli.StringFlag{
			Name:   "gemspec",
			Usage:  "path to gemspec",
			EnvVar: "PLUGIN_GEMSPEC",
		},
		cli.StringFlag{
			Name:   "host",
			Usage:  "rubygems host",
			EnvVar: "RUBYGEMS_HOST,PLUGIN_HOST",
		},
		cli.StringFlag{
			Name:   "gemname",
			Usage:  "Name of the gem",
			EnvVar: "PLUGIN_GEMNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "password (only required without api_key)",
			EnvVar: "RUBYGEMS_PASSWORD,PLUGIN_PASSWORD",
		},
		cli.StringFlag{
			Name:   "username",
			Usage:  "username (only required without api_key)",
			EnvVar: "RUBYGEMS_USERNAME,PLUGIN_USERNAME",
		},
		cli.BoolFlag{
			Name:   "skip_cleanup",
			Usage:  "flag to deploy from the current file state",
			EnvVar: "PLUGIN_SKIP_CLEANUP",
		},

		//
		// repo args
		//

		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Name:    c.String("repo.name"),
		},
		Config: Config{
			ApiKey:      c.String("api_key"),
			Gemspec:     c.String("gemspec"),
			Host:        c.String("host"),
			Gemname:     c.String("gemname"),
			Password:    c.String("password"),
			Username:    c.String("username"),
			SkipCleanup: c.Bool("skip_cleanup"),
		},
	}

	return plugin.Exec()
}
