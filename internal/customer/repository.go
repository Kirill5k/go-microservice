package customer

import (
	"context"
	"kirill5k/go/microservice/internal/database"
)

type Repository interface {
	FindBy(ctx context.Context, email string) ([]Customer, error)
}

type PostgresRepository struct {
	client *database.PostgresClient
}

func (pr *PostgresRepository) FindBy(ctx context.Context, email string) ([]Customer, error) {
	var customers []Customer
	result := pr.client.DB.WithContext(ctx).Where(Customer{Email: email}).Find(&customers)
	return customers, result.Error
}

func NewPostgresRepository(client *database.PostgresClient) *PostgresRepository {
	return &PostgresRepository{client}
}
