package ssh

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ONSdigital/dp/cmd/ansible"
	"github.com/ONSdigital/dp/cmd/aws"
	"github.com/ONSdigital/dp/cmd/config"
	"github.com/ONSdigital/dp/cmd/launch"
	"gopkg.in/urfave/cli.v1"
)

// Command returns an instance of the ssh command
func Command(cfg config.Config) cli.Command {
	cmd := cli.Command{
		Name:  "ssh",
		Usage: "access an environment using ssh",
	}

	for _, e := range cfg.Environments {
		var env = e
		subcommand := cli.Command{
			Name:  env.Name,
			Usage: "ssh to " + env.Name,
		}

		groups, err := ansible.GetGroupsForEnvironment(cfg, env.Name)
		if err != nil {
			fmt.Printf("error loading ansible hosts for %s: %s\n", env.Name, err)
			continue
		}

		for _, g := range groups {
			var grp = g
			groupcmd := cli.Command{
				Name:  grp,
				Usage: "ssh to " + grp + " in " + env.Name,
				Action: func(c *cli.Context) error {
					if len(cfg.SSHUser) == 0 {
						return fmt.Errorf("DP_SSH_USER environment variable must be set")
					}

					idx := c.Args().First()
					rIndex := int(-1)

					if len(idx) > 0 {
						idxInt, err := strconv.Atoi(idx)
						if err != nil {
							return fmt.Errorf("invalid numeric value for index: %s", idx)
						}
						rIndex = idxInt
					}

					fmt.Println("ssh to " + grp + " in " + env.Name)
					r, err := aws.ListEC2ByAnsibleGroup(env.Name, grp)
					if err != nil {
						return fmt.Errorf("error fetching ec2: %s", err)
					}
					if len(r) == 0 {
						return errors.New("no matching instances found")
					}
					var inst aws.EC2Result
					if len(r) > 1 {
						if rIndex < 0 {
							for i, v := range r {
								fmt.Printf("[%d] %s: %s %s\n", i, v.Name, v.IPAddress, v.AnsibleGroups)
							}
							return errors.New("use an index to select a specific instance")
						}

						inst = r[rIndex]
					} else {
						inst = r[0]
					}
					fmt.Printf("[%d] %s: %s %s\n", rIndex, inst.Name, inst.IPAddress, inst.AnsibleGroups)
					return launch.SSH(cfg, cfg.SSHUser, inst.IPAddress)
				},
			}
			subcommand.Subcommands = append(subcommand.Subcommands, groupcmd)
		}

		cmd.Subcommands = append(cmd.Subcommands, subcommand)
	}

	return cmd
}
