package customer

import (
	"context"
	"kirill5k/go/microservice/internal/common"
	"kirill5k/go/microservice/internal/database"
)

type Repository interface {
	FindBy(ctx context.Context, email string) ([]Customer, error)
}

type PostgresRepository struct {
	client *database.PostgresClient
}

func NewPostgresRepository(client *database.PostgresClient) *PostgresRepository {
	return &PostgresRepository{client}
}

func toDomain(e Entity) Customer {
	return Customer{
		ID:        e.ID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Email:     e.Email,
		Phone:     e.Phone,
		Address:   e.Address,
	}
}

func (pr *PostgresRepository) FindBy(ctx context.Context, email string) ([]Customer, error) {
	var entities []Entity
	result := pr.client.DB.WithContext(ctx).Where(Entity{Email: email}).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return common.Map(entities, toDomain), nil
}
