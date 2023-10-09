package customer

import "context"

type Service struct {
	repository Repository
}

func (svc *Service) FindBy(ctx context.Context, email string) ([]Customer, error) {
	return svc.repository.FindBy(ctx, email)
}

func (svc *Service) Create(ctx context.Context, newCust NewCustomer) (Customer, error) {
	return svc.repository.Create(ctx, newCust)
}

func NewService(repository Repository) *Service {
	return &Service{repository}
}
