package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OptionsValuesServiceInterface ...
type OptionsValuesServiceInterface interface {
	Create(context.Context, *models.OptionValue) error
	FindAll(context.Context, int, int) ([]*models.OptionValue, error)
	Find(context.Context, int) (*models.OptionValue, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.OptionValue) error
}

// OptionsValuesService ...
type OptionsValuesService struct {
	service *Service
}

// Create ...
func (s *OptionsValuesService) Create(ctx context.Context, u *models.OptionValue) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}

// FindAll ...
func (s *OptionsValuesService) FindAll(ctx context.Context, offset, limit int) ([]*models.OptionValue, error) {
	return nil, nil
}

// Find ...
func (s *OptionsValuesService) Find(ctx context.Context, id int) (*models.OptionValue, error) {
	return nil, nil
}

// Delete ...
func (s *OptionsValuesService) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (s *OptionsValuesService) Update(ctx context.Context, id int, u *models.OptionValue) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}
