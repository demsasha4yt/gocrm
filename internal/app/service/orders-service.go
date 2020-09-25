package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OrdersServiceInterface ...
type OrdersServiceInterface interface {
	Create(context.Context, *models.Order) error
	FindAll(context.Context, int, int) ([]*models.Order, error)
	Find(context.Context, int) (*models.Order, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Order) error
}

// OrdersService ...
type OrdersService struct {
	service *Service
}

// Create ...
func (s *OrdersService) Create(ctx context.Context, u *models.Order) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}

// FindAll ...
func (s *OrdersService) FindAll(ctx context.Context, offset, limit int) ([]*models.Order, error) {
	return nil, nil
}

// Find ...
func (s *OrdersService) Find(ctx context.Context, id int) (*models.Order, error) {
	return nil, nil
}

// Delete ...
func (s *OrdersService) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (s *OrdersService) Update(ctx context.Context, id int, u *models.Order) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}
