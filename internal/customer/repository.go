package customer

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kirill5k/go/microservice/internal/common"
	"kirill5k/go/microservice/internal/database"
)

type Repository interface {
	Get(ctx context.Context, id uuid.UUID) (*Customer, error)
	FindBy(ctx context.Context, email string) ([]Customer, error)
	Create(ctx context.Context, customer *NewCustomer) (*Customer, error)
	Update(ctx context.Context, customer *Customer) (*Customer, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type postgresRepository struct {
	client *database.PostgresClient
}

func NewPostgresRepository(client *database.PostgresClient) Repository {
	return &postgresRepository{client}
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

func (c *Customer) toEntity() *customer {
	return &customer{
		ID:        c.ID,
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

func (pr *postgresRepository) FindBy(ctx context.Context, email string) ([]Customer, error) {
	var entities []customer
	result := pr.client.DB.WithContext(ctx).Where(customer{Email: email}).Find(&entities)
	if result.Error != nil {
		return nil, handleError(result.Error)
	}
	return common.Map(entities, toDomain), nil
}

func (pr *postgresRepository) Create(ctx context.Context, newCust *NewCustomer) (*Customer, error) {
	entity := newCust.toEntity()
	result := pr.client.DB.WithContext(ctx).Create(&entity)
	if result.Error != nil {
		return nil, handleError(result.Error)
	}

	return entity.toDomain(), nil
}

func (pr *postgresRepository) Get(ctx context.Context, id uuid.UUID) (*Customer, error) {
	var entity *customer
	result := pr.client.DB.WithContext(ctx).Where(customer{ID: id}).First(&entity)
	if result.Error == nil {
		return entity.toDomain(), nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &common.NotFoundError{ID: id, Entity: "customer"}
	}

	return nil, handleError(result.Error)
}

func (pr *postgresRepository) Update(ctx context.Context, cust *Customer) (*Customer, error) {
	result := pr.client.DB.WithContext(ctx).Clauses(clause.Returning{}).Where(customer{ID: cust.ID}).Updates(cust.toEntity())

	if result.Error != nil {
		return nil, handleError(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, &common.NotFoundError{ID: cust.ID, Entity: "customer"}
	}

	return cust, nil
}

func (pr *postgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := pr.client.DB.WithContext(ctx).Delete(customer{ID: id})

	if result.Error != nil {
		return handleError(result.Error)
	}

	if result.RowsAffected == 0 {
		return &common.NotFoundError{ID: id, Entity: "customer"}
	}

	return nil
}

func handleError(err error) error {
	switch e := err.(type) {
	case *pgconn.PgError:
		if e.Code == "23505" {
			return &common.ConflictError{Detail: e.Detail}
		}
		return errors.New(e.Message)
	default:
		if errors.Is(e, gorm.ErrDuplicatedKey) {
			return &common.ConflictError{Detail: e.Error()}
		}
		return e
	}

}
