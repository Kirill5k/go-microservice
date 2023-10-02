package database

type Config struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	SSLMode  bool
}

func DefaultPostgresConfig() Config {
	return Config{
		Host:     "localhost",
		Port:     5432,
		DBName:   "postgres",
		User:     "postgres",
		Password: "postgres",
		SSLMode:  false,
	}
}
