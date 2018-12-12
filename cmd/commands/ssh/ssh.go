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

		groups, err := ansible.GetGroupsForEnvironment(cfg, env.Name)
		if err != nil {
			fmt.Printf("error loading ansible hosts for %s: %s\n", env.Name, err)
			continue
		}

		var actionFunc = func(grp string, env config.Environment, scp bool) func(c *cli.Context) error {
			return func(c *cli.Context) error {
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

				// if scp {
				// 	fmt.Println("scp to " + grp + " in " + env.Name)
				// } else {
				fmt.Println("ssh to " + grp + " in " + env.Name)
				// }

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
				// if scp {
				// 	return launch.SCP(cfg, cfg.SSHUser, inst.IPAddress, c.Args().Tail()...)
				// } else {
				return launch.SSH(cfg, cfg.SSHUser, inst.IPAddress)
				// }
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
