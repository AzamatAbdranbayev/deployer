package config

type Config struct {
	Server *ServerConfig `json:"server"`
}

type ServerConfig struct {
	Addr    string `json:"addr" env:"SERVER_ADDR"`
	Name    string `json:"name"`
	Version string `json:"version" env:"VERSION"`
}

func NewConfig(name, version string) *Config {
	return &Config{
		Server: &ServerConfig{
			Name:    name,
			Version: version,
		},
	}
}
