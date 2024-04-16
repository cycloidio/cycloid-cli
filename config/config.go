package config

import (
	"fmt"
	"io/ioutil"

	"github.com/adrg/xdg"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var (
	appName = "cycloid-cli"
	path    = "config.yaml"
)

// Config is the structure handling the config
// of the CLI
type Config struct {
	// Organizations is the list of Organization where the user
	// is currently logged in
	Organizations map[string]Organization `yaml:"organizations"`
}

// Organization is an organization where the user
// is logged in
type Organization struct {
	// Organization token
	Token string `yaml:"token"`
}

// Read will read the config from the
// path and returns a config struct
func Read() (*Config, error) {
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", appName, path))
	if err != nil {
		return &Config{
			Organizations: make(map[string]Organization),
		}, errors.Wrap(err, "unable to find XDG config path")
	}
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		// we return an empty Config in case it's the first time we try to access
		// the config and it does not exist yet
		return &Config{
			Organizations: make(map[string]Organization),
		}, errors.Wrap(err, "unable to read config from file")
	}
	var c Config
	if err := yaml.Unmarshal(content, &c); err != nil {
		return nil, errors.Wrap(err, "unable to decode config from file")
	}
	return &c, nil
}

// Write will write the config into the
// path location
func Write(c *Config) error {
	content, err := yaml.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "unable to marshal config structure")
	}
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", appName, path))
	if err != nil {
		return errors.Wrap(err, "unable to find XDG config path")
	}

	return ioutil.WriteFile(configFilePath, content, 0644)
}
