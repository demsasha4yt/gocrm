package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// UnitsServiceInterface ...
type UnitsServiceInterface interface {
	Create(context.Context, *models.Unit) error
	Find(context.Context, int) (*models.Unit, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Unit) error
}

// UnitsService ...
type UnitsService struct {
	service *Service
}

// Create ..
func (s *UnitsService) Create(ctx context.Context, u *models.Unit) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.service.store.Units().Create(ctx, u)
}

// Find ...
func (s *UnitsService) Find(ctx context.Context, id int) (*models.Unit, error) {
	return s.service.store.Units().Find(ctx, id)
}

// Delete ..
func (s *UnitsService) Delete(ctx context.Context, id int) error {
	return s.service.store.Units().Delete(ctx, id)
}

// Update ...
func (s *UnitsService) Update(ctx context.Context, id int, u *models.Unit) error {
	if err := u.Validate(); err != nil {
		return err
	}

	return s.service.store.Units().Update(ctx, id, u)
}
