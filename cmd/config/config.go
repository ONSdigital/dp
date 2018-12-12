package config

import (
	"go/build"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is the application config
type Config struct {
	CommonApps      []App         `yaml:"common_apps"`
	EnvironmentApps []App         `yaml:"environment_apps"`
	Environments    []Environment `yaml:"environments"`
	SetupRepo       string        `yaml:"setupRepo"`

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

	if _, err := os.Stat(path); err == nil {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return config, err
		}

		err = yaml.Unmarshal(b, &config)
		if err != nil {
			return config, err
		}
	}

	return config, nil
}
