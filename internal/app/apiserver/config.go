package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"build_addr"`
	LogLevel string `toml:"log_level`
}

func NewConfig() *Config {

	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
