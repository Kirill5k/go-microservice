package customer

import "context"

type Service struct {
	repository Repository
}

func (svc *Service) FindBy(ctx context.Context, email string) ([]Customer, error) {
	return svc.repository.FindBy(ctx, email)
}

func NewService(repository Repository) Service {
	return Service{repository}
}
