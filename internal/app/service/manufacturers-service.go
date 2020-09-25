package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

//ManufacturersServiceInterface ...
type ManufacturersServiceInterface interface {
	Create(context.Context, *models.Manufacturer) error
	FindAll(context.Context, int, int) ([]*models.Manufacturer, error)
	Find(context.Context, int) (*models.Manufacturer, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Manufacturer) error
}

// ManufacturersService ...
type ManufacturersService struct {
	service *Service
}

// Create ...
func (s *ManufacturersService) Create(ctx context.Context, u *models.Manufacturer) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.service.store.Manufacturers().Create(ctx, u)
}

// FindAll ...
func (s *ManufacturersService) FindAll(ctx context.Context, offset, limit int) ([]*models.Manufacturer, error) {
	return s.service.store.Manufacturers().FindAll(ctx, offset, limit)
}

// Find ...
func (s *ManufacturersService) Find(ctx context.Context, id int) (*models.Manufacturer, error) {
	return s.service.store.Manufacturers().Find(ctx, id)
}

// Delete ...
func (s *ManufacturersService) Delete(ctx context.Context, id int) error {
	return s.service.store.Manufacturers().Delete(ctx, id)
}

// Update ...
func (s *ManufacturersService) Update(ctx context.Context, id int, u *models.Manufacturer) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.service.store.Manufacturers().Update(ctx, id, u)
}
