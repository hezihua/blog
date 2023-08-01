package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	c *Config
)

func C() *Config {
	if c == nil {
		panic("config not initialized")
	}
	return c
}

func LoadConfigFromToml(filepath string) error {
	c = NewConfig()
	_, err := toml.DecodeFile(filepath, c)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfigFromEnv() error {
	c = NewConfig()
	env.Parse(c)
	return nil
}