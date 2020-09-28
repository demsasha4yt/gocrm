package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// VariationsRepository ...
type VariationsRepository struct {
	store *Store
}

// Create ...
func (r *VariationsRepository) Create(ctx context.Context, u *models.Variation) error {
	return nil
}

// Find ...
func (r *VariationsRepository) Find(ctx context.Context, id int) (*models.Variation, error) {
	return nil, nil
}

// FindAll ...
func (r *VariationsRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.Variation, error) {
	return nil, nil
}

// Delete ...
func (r *VariationsRepository) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (r *VariationsRepository) Update(ctx context.Context, id int, u *models.Variation) error {
	return nil
}
