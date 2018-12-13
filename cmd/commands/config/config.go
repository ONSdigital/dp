package config

import (
	"fmt"
	"path/filepath"

	"github.com/ONSdigital/dp/cmd/config"
	"github.com/ONSdigital/dp/cmd/launch"
	cli "gopkg.in/urfave/cli.v1"
)

// Command returns an instance of the decrypt command
func Command(cfg config.Config) cli.Command {
	cmd := cli.Command{
		Name:  "config",
		Usage: "decrypt an environments vault config",
	}

	for _, e := range cfg.Environments {
		var env = e
		subcommand := cli.Command{
			Name:  env.Name,
			Usage: "open " + env.Name,
			Action: func(c *cli.Context) error {
				fmt.Println("decrypting vault config for " + env.Name)
				pwd := filepath.Join(cfg.GoPath, "src", cfg.SetupRepo, "ansible")

				return launch.Command(pwd, "../scripts/decrypt_inline_vault", fmt.Sprintf("%s@.%s.pass", env.Name, env.Name), fmt.Sprintf("inventories/%s/group_vars/all", env.Name))
			},
		}
		cmd.Subcommands = append(cmd.Subcommands, subcommand)
	}

	return cmd
}
