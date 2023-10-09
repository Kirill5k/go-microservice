package customer

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kirill5k/go/microservice/internal/common"
	"kirill5k/go/microservice/internal/database"
)

type Repository interface {
	FindBy(ctx context.Context, email string) ([]Customer, error)
	Create(ctx context.Context, customer NewCustomer) (Customer, error)
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

func newCustomer(c NewCustomer) customer {
	return customer{
		ID:        uuid.NewString(),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
		Address:   c.Address,
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

func (pr *PostgresRepository) Create(ctx context.Context, newCust NewCustomer) (Customer, error) {
	var cust Customer
	entity := newCustomer(newCust)
	result := pr.client.DB.WithContext(ctx).Create(&entity)
	if result.Error == nil {
		cust = toDomain(entity)
		return cust, nil
	}

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return cust, &database.ConflictError{}
	}

	return cust, result.Error
}
