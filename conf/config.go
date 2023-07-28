package conf

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
