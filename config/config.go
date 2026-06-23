package config

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

var (
	appName = "cycloid-cli"
	path    = "config.yaml"
)

// Config handles the CLI configuration
type Config struct {
	Organizations map[string]Organization `yaml:"organizations"`
	Output        string                  `yaml:"output,omitempty"`
}

// Organization represents a logged-in organization session
type Organization struct {
	Token string `yaml:"token"`
}

func GetConfigPath() (string, error) {
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", appName, path))
	if err != nil {
		return "", fmt.Errorf("invalid config: unable to find XDG config path: %w", err)
	}

	return configFilePath, nil
}

// Read reads the config from the XDG path and returns a Config struct
func Read() (*Config, error) {
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", appName, path))
	if err != nil {
		return &Config{
			Organizations: make(map[string]Organization),
		}, fmt.Errorf("invalid config: unable to find XDG config path: %w", err)
	}
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return &Config{
			Organizations: make(map[string]Organization),
		}, fmt.Errorf("invalid config: unable to read config from file: %w", err)
	}
	var c Config
	if err := yaml.Unmarshal(content, &c); err != nil {
		return nil, fmt.Errorf("invalid config: unable to decode config from file: %w", err)
	}
	return &c, nil
}

// Write writes the config into the XDG path
func Write(c *Config) error {
	content, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("invalid config: unable to marshal config structure: %w", err)
	}
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("%s/%s", appName, path))
	if err != nil {
		return fmt.Errorf("invalid config: unable to find XDG config path: %w", err)
	}

	if err := os.WriteFile(configFilePath, content, 0o600); err != nil {
		return fmt.Errorf("invalid config: unable to write config file: %w", err)
	}
	return nil
}
