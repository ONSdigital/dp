package ansible

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/ONSdigital/dp/cmd/config"
)

// GetGroupsForEnvironment returns a list of ansible groups for the specified environment
func GetGroupsForEnvironment(cfg config.Config, environment string) ([]string, error) {
	hostsPath := filepath.Join(cfg.GoPath, "src", cfg.SetupRepo, "ansible/inventories", environment, "hosts")

	b, err := ioutil.ReadFile(hostsPath)
	if err != nil {
		return nil, err
	}

	var groups []string

	r := bufio.NewReader(bytes.NewReader(b))
	for {
		s, err := r.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.HasPrefix(s, "[") && strings.Contains(s, ":") && strings.HasSuffix(s, "]") {
			name := s[1:strings.Index(s, ":")]
			groups = append(groups, name)
		}
		if err == io.EOF {
			break
		}
	}

	return groups, nil
}
