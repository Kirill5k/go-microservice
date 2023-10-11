package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"kirill5k/go/microservice/internal/common"
	"time"
)

type Client interface {
	Ready() bool
}

type PostgresClient struct {
	DB *gorm.DB
}

func (c *PostgresClient) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	return ready == "1"
}

func NewPostgresClient(config *PostgresConfig) (*PostgresClient, error) {
	var client PostgresClient
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		common.If(config.SSLMode, "enable", "disable"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: config.TablePrefix},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return &client, err
	}
	client = PostgresClient{db}
	return &client, nil
}
