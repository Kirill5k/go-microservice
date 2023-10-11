package database

type PostgresConfig struct {
	Host        string
	Port        int
	DBName      string `mapstructure:"db-name"`
	TablePrefix string `mapstructure:"table-prefix"`
	Username    string
	Password    string
	SSLMode     bool `mapstructure:"ssl-mode"`
}

func DefaultPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		Host:        "localhost",
		Port:        5432,
		DBName:      "postgres",
		TablePrefix: "wisdom.",
		Username:    "postgres",
		Password:    "postgres",
		SSLMode:     false,
	}
}
