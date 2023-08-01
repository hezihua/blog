package conf_test

import (
	"hezihua/conf"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("./test/test.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().Auth)
}

func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().Auth)
}