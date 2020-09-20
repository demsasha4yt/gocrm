package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OptionsRepository ...
type OptionsRepository struct {
	store *Store
}

// Create ...
func (r *OptionsRepository) Create(ctx context.Context, u *models.Option) error {
	return nil
}

// Find ...
func (r *OptionsRepository) Find(ctx context.Context, id int) (*models.Option, error) {
	return nil, nil
}

// Delete ...
func (r *OptionsRepository) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (r *OptionsRepository) Update(ctx context.Context, id int, u *models.Option) error {
	return nil
}
