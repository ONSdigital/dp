package ui

import (
	"fmt"
	"strings"

	"github.com/ONSdigital/dp/cmd/config"
	"github.com/ONSdigital/dp/cmd/launch"
	"gopkg.in/urfave/cli.v1"
)

// Command returns an instance of the ssh command
func Command(cfg config.Config) cli.Command {
	cmd := cli.Command{
		Name:  "ui",
		Usage: "open an environment ui",
	}

	for _, a := range cfg.CommonApps {
		var app = a
		cmd.Subcommands = append(cmd.Subcommands, cli.Command{
			Name:  app.Name,
			Usage: "open " + app.Name,
			Action: func(c *cli.Context) error {
				fmt.Println("opening " + app.Name)
				return launch.Browser(app.URL)
			},
		})
	}

	for _, a := range cfg.EnvironmentApps {
		var app = a
		subcommand := cli.Command{
			Name:  app.Name,
			Usage: "open " + app.Name,
		}
		for _, e := range cfg.Environments {
			var env = e
			subcommand.Subcommands = append(subcommand.Subcommands, cli.Command{
				Name:  env.Name,
				Usage: "open " + app.Name + " for " + env.Name,
				Action: func(c *cli.Context) error {
					fmt.Println("opening " + app.Name + " for " + env.Name)
					return launch.Browser(strings.Replace(app.URL, "$environment", env.Name, -1))
				},
			})
		}
		cmd.Subcommands = append(cmd.Subcommands, subcommand)
	}

	return cmd
}
