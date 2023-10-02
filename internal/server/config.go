package server

type Config struct {
	Port int
}

func DefaultConfig() Config {
	return Config{8080}
}
