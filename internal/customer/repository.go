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
	Get(ctx context.Context, id uuid.UUID) (*Customer, error)
	FindBy(ctx context.Context, email string) ([]Customer, error)
	Create(ctx context.Context, customer *NewCustomer) (*Customer, error)
}

type PostgresRepository struct {
	client *database.PostgresClient
}

func NewPostgresRepository(client *database.PostgresClient) *PostgresRepository {
	return &PostgresRepository{client}
}

type customer struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
	Phone     string
	Address   string
}

func (c *NewCustomer) toEntity() *customer {
	return &customer{
		ID:        uuid.New(),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
		Address:   c.Address,
	}
}

func (c *customer) toDomain() *Customer {
	return &Customer{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
		Address:   c.Address,
	}
}

func toDomain(e *customer) Customer {
	return *e.toDomain()
}

func (pr *PostgresRepository) FindBy(ctx context.Context, email string) ([]Customer, error) {
	var entities []customer
	result := pr.client.DB.WithContext(ctx).Where(customer{Email: email}).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return common.Map(entities, toDomain), nil
}

func (pr *PostgresRepository) Create(ctx context.Context, newCust *NewCustomer) (*Customer, error) {
	entity := newCust.toEntity()
	result := pr.client.DB.WithContext(ctx).Create(&entity)
	if result.Error == nil {
		return entity.toDomain(), nil
	}

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return nil, &common.ConflictError{}
	}

	return nil, result.Error
}

func (pr *PostgresRepository) Get(ctx context.Context, id uuid.UUID) (*Customer, error) {
	var entity *customer
	result := pr.client.DB.WithContext(ctx).Where(customer{ID: id}).First(&entity)
	if result.Error == nil {
		return entity.toDomain(), nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &common.NotFoundError{ID: id, Entity: "customer"}
	}

	return nil, result.Error
}
