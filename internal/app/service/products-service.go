package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// ProductsServiceInterface ...
type ProductsServiceInterface interface {
	Create(context.Context, *models.Product) error
	FindAll(context.Context, int, int) ([]*models.Product, error)
	Find(context.Context, int) (*models.Product, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Product) error
}

// ProductsService ...
type ProductsService struct {
	service *Service
}

// Create ...
func (s *ProductsService) Create(ctx context.Context, u *models.Product) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}

// FindAll ...
func (s *ProductsService) FindAll(ctx context.Context, offset, limit int) ([]*models.Product, error) {
	return nil, nil
}

// Find ...
func (s *ProductsService) Find(ctx context.Context, id int) (*models.Product, error) {
	return nil, nil
}

// Delete ...
func (s *ProductsService) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (s *ProductsService) Update(ctx context.Context, id int, u *models.Product) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}
