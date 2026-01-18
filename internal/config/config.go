// Package config contains all the config for gator
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL    string `json:"db_url"`
	UserName string `json:"user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	p, err := getCofingFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(p)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) SetUser(name string) error {
	c.UserName = name

	if err := write(*c); err != nil {
		return err
	}

	return nil
}

func getCofingFilePath() (string, error) {
	p, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	newPath := filepath.Join(p, configFileName)

	return newPath, nil
}

func write(cfg Config) error {
	p, err := getCofingFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if err := os.WriteFile(p, data, 0644); err != nil {
		return err
	}

	return nil
}
