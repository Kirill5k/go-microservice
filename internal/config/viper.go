package config

import (
	"github.com/spf13/viper"
	"kirill5k/go/microservice/internal/database"
	"kirill5k/go/microservice/internal/server"
	"log"
)

type Config struct {
	Server   server.Config
	Postgres database.PostgresConfig
}

func LoadViperConfig() *Config {
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("failed to read viper config. %v", err)
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("failed to decode viper config into struct. %v", err)
	}
	return &c
}
