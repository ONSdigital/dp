package launch

import (
	"fmt"
	"path/filepath"

	"github.com/ONSdigital/dp/cmd/config"
)

// SSH opens an SSH connection
func SSH(cfg config.Config, user string, ip string) error {
	pwd := filepath.Join(cfg.GoPath, "src", cfg.SetupRepo, "ansible")

	return Command(pwd, "ssh", "-F", "ssh.cfg", fmt.Sprintf("%s@%s", user, ip))
}
