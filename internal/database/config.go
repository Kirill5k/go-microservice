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
