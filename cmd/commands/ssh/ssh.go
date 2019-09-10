package ssh

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ONSdigital/dp/cmd/ansible"
	"github.com/ONSdigital/dp/cmd/aws"
	"github.com/ONSdigital/dp/cmd/config"
	"github.com/ONSdigital/dp/cmd/launch"
	cli "gopkg.in/urfave/cli.v1"
)

const (
	colWarn   = "\033[31;1m"
	colAlt    = "\033[32m"
	colHi     = "\033[1m"
	colReset  = "\033[0m"
	widthName = -30
	widthIP   = -15
)

// Command returns an instance of the ssh command
func Command(cfg config.Config) []cli.Command {
	sshCmd := cli.Command{
		Name:  "ssh",
		Usage: "access an environment using ssh",
	}
	// scpCmd := cli.Command{
	// 	Name:  "scp",
	// 	Usage: "access an environment using scp",
	// }

	for _, e := range cfg.Environments {
		var env = e
		sshSubcommand := cli.Command{
			Name:  env.Name,
			Usage: "ssh to " + env.Name,
		}
		// scpSubcommand := cli.Command{
		// 	Name:  env.Name,
		// 	Usage: "scp to " + env.Name,
		// }

		profile := cfg.DefaultEnvProfile
		if e.Profile != "" {
			profile = e.Profile
		}

		groups, err := ansible.GetGroupsForEnvironment(cfg, env.Name)
		if err != nil {
			fmt.Printf("error loading ansible hosts for %s: %s\n", env.Name, err)
			continue
		}

		var actionFunc = func(grp string, env config.Environment, scp bool) func(c *cli.Context) error {
			return func(c *cli.Context) error {
				if len(cfg.SSHUser) == 0 {
					return cli.NewExitError("DP_SSH_USER environment variable must be set", 22)
				}

				colEnv := colAlt
				if env.Name == "production" {
					colEnv = colWarn
				}
				fmt.Println("ssh to " + colAlt + grp + colReset + " in " + colEnv + env.Name + colReset)

				r, err := aws.ListEC2ByAnsibleGroup(env.Name, profile, grp)
				if err != nil {
					return fmt.Errorf("error fetching ec2: %s", err)
				}
				if len(r) == 0 {
					return errors.New("no matching instances found")
				}

				var itemChosen int
				if argAsString := c.Args().First(); len(argAsString) > 0 {
					var err error
					itemChosen, err = strconv.Atoi(argAsString)
					if err != nil || itemChosen == 0 || itemChosen > len(r) {
						return fmt.Errorf("invalid numeric value: %q - must be 1..%d", argAsString, len(r))
					}
				}

				for i, v := range r {
					if itemChosen > 0 && itemChosen != i+1 {
						continue // args selected one, but not this one
					}
					colSwitch := ""
					if i%2 == 0 {
						colSwitch = colAlt
					}
					fmt.Printf(colSwitch+"["+colHi+"%2d"+colReset+colSwitch+"] %*s: "+colHi+"%*s"+colReset+" "+colSwitch+"%s"+colReset+"\n",
						i+1,
						widthName, v.Name,
						widthIP, v.IPAddress,
						v.AnsibleGroups,
					)
				}
				if itemChosen == 0 {
					return cli.NewExitError("use an number to select a specific instance", 2)
				}

				return launch.SSH(cfg, cfg.SSHUser, r[itemChosen-1].IPAddress)
			}
		}

		for _, g := range groups {
			var grp = g
			sshGroupCmd := cli.Command{
				Name:   grp,
				Usage:  "ssh to " + grp + " in " + env.Name,
				Action: actionFunc(grp, env, false),
			}
			// scpGroupCmd := cli.Command{
			// 	Name:   grp,
			// 	Usage:  "scp to " + grp + " in " + env.Name,
			// 	Action: actionFunc(grp, env, true),
			// }
			sshSubcommand.Subcommands = append(sshSubcommand.Subcommands, sshGroupCmd)
			// scpSubcommand.Subcommands = append(scpSubcommand.Subcommands, scpGroupCmd)
		}

		sshCmd.Subcommands = append(sshCmd.Subcommands, sshSubcommand)
		// scpCmd.Subcommands = append(scpCmd.Subcommands, scpSubcommand)
	}

	// return []cli.Command{sshCmd, scpCmd}
	return []cli.Command{sshCmd}
}
