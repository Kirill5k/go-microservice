package database

type DatabaseConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	SSLMode  bool
}

func DefaultDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		DBName:   "postgres",
		User:     "postgres",
		Password: "postgres",
		SSLMode:  false,
	}
}
