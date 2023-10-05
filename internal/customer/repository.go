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
	var entities []Entity
	result := pr.client.DB.WithContext(ctx).Where(Entity{Email: email}).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	customers := make([]Customer, len(entities))
	for i, entity := range entities {
		customers[i] = entity.toDomain()
	}
	return customers, result.Error
}

func NewPostgresRepository(client *database.PostgresClient) *PostgresRepository {
	return &PostgresRepository{client}
}
