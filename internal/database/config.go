package database

type PostgresConfig struct {
	Host        string
	Port        int
	DBName      string
	TablePrefix string
	User        string
	Password    string
	SSLMode     bool
}

func DefaultPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		Host:        "localhost",
		Port:        5432,
		DBName:      "postgres",
		TablePrefix: "wisdom.",
		User:        "postgres",
		Password:    "postgres",
		SSLMode:     false,
	}
}
