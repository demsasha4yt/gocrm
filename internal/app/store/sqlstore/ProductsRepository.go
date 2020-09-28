package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// ProductsRepository ...
type ProductsRepository struct {
	store *Store
}

// Create ...
func (r *ProductsRepository) Create(ctx context.Context, u *models.Product) error {
	return nil
}

// Find ...
func (r *ProductsRepository) Find(ctx context.Context, id int) (*models.Product, error) {
	return nil, nil
}

// FindAll ...
func (r *ProductsRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.Product, error) {
	return nil, nil
}

// Delete ...
func (r *ProductsRepository) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (r *ProductsRepository) Update(ctx context.Context, id int, u *models.Product) error {
	return nil
}
