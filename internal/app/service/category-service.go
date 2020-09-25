package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// CategoriesServiceInterface ...
type CategoriesServiceInterface interface {
	Create(context.Context, *models.Category) error
	FindAll(context.Context, int, int) ([]*models.Category, error)
	Find(context.Context, int) (*models.Category, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Category) error
}

// CategoriesService ...
type CategoriesService struct {
	service *Service
}

// Create ...
func (s *CategoriesService) Create(ctx context.Context, u *models.Category) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.service.store.Categories().Create(ctx, u)
}

// FindAll ...
func (s *CategoriesService) FindAll(ctx context.Context, offset, limit int) ([]*models.Category, error) {
	return s.service.store.Categories().FindAll(ctx, offset, limit)
}

// Find ...
func (s *CategoriesService) Find(ctx context.Context, id int) (*models.Category, error) {
	return s.service.store.Categories().Find(ctx, id)
}

// Delete ...
func (s *CategoriesService) Delete(ctx context.Context, id int) error {
	return s.service.store.Categories().Delete(ctx, id)
}

// Update ...
func (s *CategoriesService) Update(ctx context.Context, id int, u *models.Category) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.service.store.Categories().Update(ctx, id, u)
}
