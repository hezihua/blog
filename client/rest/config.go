package rest

func NewDefaultConfig() *Config {
	return &Config{
		Url:      "http://127.0.0.1:8050/vblog/api/v1",
		Username: "adming",
		Password: "123456",
	}
}

type Config struct {
	Url      string
	Username string
	Password string
}