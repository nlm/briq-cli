package config

import "errors"

var ConfigFile Config

type Config struct {
	SecretKey string              `mapstructure:"secret_key"`
	Groups    map[string][]string `mapstructure:"groups"`
}

func (c Config) Check() error {
	if c.SecretKey == "" {
		return errors.New("secret_key must be defined")
	}
	return nil
}
