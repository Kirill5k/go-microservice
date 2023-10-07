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

type customer struct {
	ID        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
	Phone     string
	Address   string
}

func toDomain(e customer) Customer {
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
	var entities []customer
	result := pr.client.DB.WithContext(ctx).Where(customer{Email: email}).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return common.Map(entities, toDomain), nil
}
