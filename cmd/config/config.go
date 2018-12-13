package config

import (
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Config is the application config
type Config struct {
	CommonApps      []App         `yaml:"common_apps"`
	EnvironmentApps []App         `yaml:"environment_apps"`
	Environments    []Environment `yaml:"environments"`

	SetupRepo string `yaml:"setupRepo"`

	GoPath  string `yaml:"-"`
	SSHUser string `yaml:"-"`
}

// App represents an app
type App struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

// Environment represents an environment
type Environment struct {
	Name string `yaml:"name"`
}

// GetDPSetupPath resolves the location of the dp-setup repo
func (c Config) GetDPSetupPath() string {
	// If it's an absolute path, just return it
	if len(c.SetupRepo) > 0 && c.SetupRepo[0] == '/' {
		return c.SetupRepo
	}
	// Otherwise assume it's in GOPATH/src
	return filepath.Join(c.GoPath, "src", c.SetupRepo)
}

// Load loads the config
func Load() (Config, error) {
	var config Config

	config.GoPath = os.Getenv("GOPATH")
	if config.GoPath == "" {
		config.GoPath = build.Default.GOPATH
	}

	config.SSHUser = os.Getenv("DP_SSH_USER")

	path := os.Getenv("DP_CONFIG")
	if len(path) == 0 {
		path = "config.yml"
	}

	defer func() {
		// Do this last so it overrides anything loaded in the config file
		setupRepo := os.Getenv("DP_SETUP")
		if len(setupRepo) > 0 {
			config.SetupRepo = setupRepo
		}
	}()

	var b []byte
	var err error

	if strings.HasPrefix(path, "https://") {
		res, err := httpClient.Get(path)
		if err != nil {
			return config, err
		}

		defer func() {
			io.Copy(ioutil.Discard, res.Body)
		}()

		if res.StatusCode >= 400 {
			return config, fmt.Errorf("unexpected status code fetching config: %d", res.StatusCode)
		}

		b, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return config, err
		}
	} else {
		if _, err := os.Stat(path); err != nil {
			if path == "config.yml" {
				return config, nil
			}
			return config, err
		}

		b, err = ioutil.ReadFile(path)
		if err != nil {
			return config, err
		}
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
