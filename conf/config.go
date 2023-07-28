package conf

import "fmt"

func NewConfig() *Config {
	return &Config{
		Auth: newAuth(),
	}
}

type Config struct {
	Auth *Auth `toml:"auth"`
}

func newAuth() *Auth {
	return &Auth{
		Username: "admin",
		Password: "123456",
	}
}

type Auth struct {
	Username string `toml:"username" env:"AUTH_USERNAME"`
	Password string `toml:"password" env:"AUTH_PASSWORD"`
}

func (a *Auth) Validate(username string, password string) error {
	if !(a.Username == username && a.Password == password) {
		return fmt.Errorf("username or password error")
	}
	return nil
}
