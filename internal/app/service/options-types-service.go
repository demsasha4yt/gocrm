package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OptionsTypesServiceInterface ...
type OptionsTypesServiceInterface interface {
	Create(context.Context, *models.OptionType) error
	FindAll(context.Context, int, int) ([]*models.OptionType, error)
	Find(context.Context, int) (*models.OptionType, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.OptionType) error
}

// OptionsTypesService ...
type OptionsTypesService struct {
	service *Service
}

// Create ...
func (s *OptionsTypesService) Create(ctx context.Context, u *models.OptionType) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}

// FindAll ...
func (s *OptionsTypesService) FindAll(ctx context.Context, offset, limit int) ([]*models.OptionType, error) {
	return nil, nil
}

// Find ...
func (s *OptionsTypesService) Find(ctx context.Context, id int) (*models.OptionType, error) {
	return nil, nil
}

// Delete ...
func (s *OptionsTypesService) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (s *OptionsTypesService) Update(ctx context.Context, id int, u *models.OptionType) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}
