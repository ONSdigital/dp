package remote

import (
	"log"

	"github.com/ONSdigital/dp/cmd/aws"
	"github.com/ONSdigital/dp/cmd/config"
	cli "gopkg.in/urfave/cli.v1"
)

// Command returns an instance of the remote command
func Command(cfg config.Config) cli.Command {
	cmd := cli.Command{
		Name:  "remote",
		Usage: "allow or deny remote access to an environment",
	}

	allowCmd := cli.Command{
		Name:  "allow",
		Usage: "allow access to environment",
	}
	denyCmd := cli.Command{
		Name:  "deny",
		Usage: "deny access to environment",
	}
	concourseCmd := cli.Command{
		Name:  "concourse",
		Usage: "control access to concourse",
		Subcommands: []cli.Command{
			{
				Name:  "allow",
				Usage: "allow access to concourse",
				Action: func(c *cli.Context) error {
					log.Println("allow access to concourse")
					return aws.AllowIPForConcourse(cfg)
				},
			},
			{
				Name:  "deny",
				Usage: "deny access to concourse",
				Action: func(c *cli.Context) error {
					log.Println("deny access to concourse")
					return aws.DenyIPForConcourse(cfg)
				},
			},
		},
	}
	for _, e := range cfg.Environments {
		var env = e
		profile := cfg.DefaultEnvProfile
		if e.Profile != "" {
			profile = e.Profile
		}
		allowCmd.Subcommands = append(allowCmd.Subcommands, cli.Command{
			Name:  env.Name,
			Usage: "allow access to " + env.Name,
			Action: func(c *cli.Context) error {
				log.Println("allow access to " + env.Name)
				return aws.AllowIPForEnvironment(cfg, env.Name, profile)
			},
		})
		denyCmd.Subcommands = append(denyCmd.Subcommands, cli.Command{
			Name:  env.Name,
			Usage: "denying access to " + env.Name,
			Action: func(c *cli.Context) error {
				log.Println("deny access to " + env.Name)
				return aws.DenyIPForEnvironment(cfg, env.Name, profile)
			},
		})
	}

	cmd.Subcommands = append(cmd.Subcommands, allowCmd, denyCmd, concourseCmd)

	return cmd
}
