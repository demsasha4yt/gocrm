package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OptionsServiceInterface ...
type OptionsServiceInterface interface {
	Create(context.Context, *models.Option) error
	FindAll(context.Context, int, int) ([]*models.Option, error)
	Find(context.Context, int) (*models.Option, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Option) error
}

// OptionsService ...
type OptionsService struct {
	service *Service
}

// Create ...
func (s *OptionsService) Create(ctx context.Context, u *models.Option) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}

// FindAll ...
func (s *OptionsService) FindAll(ctx context.Context, offset, limit int) ([]*models.Option, error) {
	return nil, nil
}

// Find ...
func (s *OptionsService) Find(ctx context.Context, id int) (*models.Option, error) {
	return nil, nil
}

// Delete ...
func (s *OptionsService) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (s *OptionsService) Update(ctx context.Context, id int, u *models.Option) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}
