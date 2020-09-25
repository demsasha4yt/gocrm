package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// VariationsServiceInterface ...
type VariationsServiceInterface interface {
	Create(context.Context, *models.Variation) error
	FindAll(context.Context, int, int) ([]*models.Variation, error)
	Find(context.Context, int) (*models.Variation, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Variation) error
}

// VariationsService ...
type VariationsService struct {
	service *Service
}

// Create ...
func (s *VariationsService) Create(ctx context.Context, u *models.Variation) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}

// FindAll ...
func (s *VariationsService) FindAll(ctx context.Context, offset, limit int) ([]*models.Variation, error) {
	return nil, nil
}

// Find ...
func (s *VariationsService) Find(ctx context.Context, id int) (*models.Variation, error) {
	return nil, nil
}

// Delete ...
func (s *VariationsService) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (s *VariationsService) Update(ctx context.Context, id int, u *models.Variation) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}
