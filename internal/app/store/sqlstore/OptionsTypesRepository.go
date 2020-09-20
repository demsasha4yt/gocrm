package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OptionsTypesRepository struct { ...
type OptionsTypesRepository struct {
	store *Store
}

// Create ...
func (r *OptionsTypesRepository) Create(ctx context.Context, u *models.OptionType) error {
	return nil
}

// Find ...
func (r *OptionsTypesRepository) Find(ctx context.Context, id int) (*models.OptionType, error) {
	return nil, nil
}

// Delete ...
func (r *OptionsTypesRepository) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (r *OptionsTypesRepository) Update(ctx context.Context, id int, u *models.OptionType) error {
	return nil
}
