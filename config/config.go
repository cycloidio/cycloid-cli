package config

import (
	"fmt"
	"io/ioutil"

	"github.com/adrg/xdg"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	appName = "cycloid-cli"
	path    = "config.yaml"
)

// Config is the structure handling the config
// of the CLI
type Config struct {
	// Token can be, at the moment, the user token
	// or the organization token
	Token string `yaml:"token"`
}

// ReadConfig will read the config from the
// path and returns a config struct
func ReadConfig() (*Config, error) {
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", appName, path))
	if err != nil {
		return nil, errors.Wrap(err, "unable to find XDG config path")
	}
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read config from file")
	}
	var c Config
	if err := yaml.Unmarshal(content, &c); err != nil {
		return nil, errors.Wrap(err, "unable to decode config from file")
	}
	return &c, nil
}

// WriteConfig will write the config into the
// path location
func WriteConfig(c *Config) error {
	content, err := yaml.Marshal(c)
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", appName, path))
	if err != nil {
		return errors.Wrap(err, "unable to find XDG config path")
	}
	if err != nil {
		return errors.Wrap(err, "unable to marshal config structure")
	}
	return ioutil.WriteFile(configFilePath, content, 0644)
}
