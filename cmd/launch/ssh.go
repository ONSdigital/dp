package launch

import (
	"fmt"
	"path/filepath"

	"github.com/ONSdigital/dp/cmd/config"
)

// SSH opens an SSH connection
func SSH(cfg config.Config, user string, ip string) error {
	pwd := filepath.Join(cfg.GetDPSetupPath(), "ansible")

	return Command(pwd, "ssh", "-F", "ssh.cfg", fmt.Sprintf("%s@%s", user, ip))
}

// SCP opens an SCP connection
func SCP(cfg config.Config, user string, ip string, args ...string) error {
	pwd := filepath.Join(cfg.GetDPSetupPath(), "ansible")

	newArgs := []string{"-F", "ssh.cfg", fmt.Sprintf("%s@%s", user, ip)}
	newArgs = append(newArgs, args...)

	return Command(pwd, "ssh", newArgs...)
}
