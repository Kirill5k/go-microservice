package customer

import (
	"context"
	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func (svc *Service) FindBy(ctx context.Context, email string) ([]Customer, error) {
	return svc.repository.FindBy(ctx, email)
}

func (svc *Service) Create(ctx context.Context, newCust *NewCustomer) (*Customer, error) {
	return svc.repository.Create(ctx, newCust)
}

func (svc *Service) Get(ctx context.Context, id uuid.UUID) (*Customer, error) {
	return svc.repository.Get(ctx, id)
}

func (svc *Service) Update(ctx context.Context, cust *Customer) (*Customer, error) {
	return svc.repository.Update(ctx, cust)
}

func NewService(repository Repository) *Service {
	return &Service{repository}
}
