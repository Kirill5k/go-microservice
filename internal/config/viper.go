package config

import (
	"fmt"
	"github.com/spf13/viper"
	"kirill5k/go/microservice/internal/database"
	"kirill5k/go/microservice/internal/server"
	"log"
	"os"
	"strings"
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

	for _, k := range v.AllKeys() {
		value := v.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			envVarName, defaultValue := getEnvVarNameWithDefaultValue(value)
			v.Set(k, getEnv(envVarName, defaultValue))
		}
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("failed to decode viper config into struct. %v", err)
	}
	return &c
}

func getEnvVarNameWithDefaultValue(stringTemplate string) (string, string) {
	envVarName := strings.TrimSuffix(strings.TrimPrefix(stringTemplate, "${"), "}")
	if strings.Contains(envVarName, ":") {
		split := strings.SplitN(envVarName, ":", 2)
		return split[0], split[1]
	}
	return envVarName, ""
}

func getEnv(envVarName, defaultValue string) string {
	value, found := os.LookupEnv(envVarName)
	if found {
		return value
	}
	if !found && defaultValue != "" {
		return defaultValue
	}
	panic(fmt.Sprintf("Missing required environment variable %s", envVarName))
}
