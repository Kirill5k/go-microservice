package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DatabaseClient interface {
	Ready() bool
}

type Client struct {
	DB *gorm.DB
}

func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	return ready == "1"
}

func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		"localhost",
		"postgres",
		"postgres",
		"postgres",
		5432,
		"disable",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "wisdom."},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{db}
	return client, nil
}
