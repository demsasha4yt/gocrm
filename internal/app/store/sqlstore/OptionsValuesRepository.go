package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OptionsValuesRepository ...
type OptionsValuesRepository struct {
	store *Store
}

// Create ...
func (r *OptionsValuesRepository) Create(ctx context.Context, u *models.OptionValue) error {
	return nil
}

// Find ...
func (r *OptionsValuesRepository) Find(ctx context.Context, id int) (*models.OptionValue, error) {
	return nil, nil
}

// Delete ...
func (r *OptionsValuesRepository) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (r *OptionsValuesRepository) Update(ctx context.Context, id int, u *models.OptionValue) error {
	return nil
}
