package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// CustomersServiceInterface ...
type CustomersServiceInterface interface {
	Create(context.Context, *models.Customer) error
	FindAll(context.Context, int, int) ([]*models.Customer, error)
	Find(context.Context, int) (*models.Customer, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Customer) error
}

// CustomersService ...
type CustomersService struct {
	service *Service
}

// Create ...
func (s *CustomersService) Create(ctx context.Context, u *models.Customer) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}

// FindAll ...
func (s *CustomersService) FindAll(ctx context.Context, offset, limit int) ([]*models.Customer, error) {
	return nil, nil
}

// Find ...
func (s *CustomersService) Find(ctx context.Context, id int) (*models.Customer, error) {
	return nil, nil
}

// Delete ...
func (s *CustomersService) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (s *CustomersService) Update(ctx context.Context, id int, u *models.Customer) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}
