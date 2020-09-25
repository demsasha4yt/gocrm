package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OptionsSoftsServiceInterface ...
type OptionsSoftsServiceInterface interface {
	Create(context.Context, *models.OptionSoft) error
	FindAll(context.Context, int, int) ([]*models.OptionSoft, error)
	Find(context.Context, int) (*models.OptionSoft, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.OptionSoft) error
}

// OptionsSoftsService ...
type OptionsSoftsService struct {
	service *Service
}

// Create ...
func (s *OptionsSoftsService) Create(ctx context.Context, u *models.OptionSoft) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}

// FindAll ...
func (s *OptionsSoftsService) FindAll(ctx context.Context, offset, limit int) ([]*models.OptionSoft, error) {
	return nil, nil
}

// Find ...
func (s *OptionsSoftsService) Find(ctx context.Context, id int) (*models.OptionSoft, error) {
	return nil, nil
}

// Delete ...
func (s *OptionsSoftsService) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (s *OptionsSoftsService) Update(ctx context.Context, id int, u *models.OptionSoft) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return nil
}
