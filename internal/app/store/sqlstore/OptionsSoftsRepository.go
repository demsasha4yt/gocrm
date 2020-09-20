package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OptionsSoftsRepository ...
type OptionsSoftsRepository struct {
	store *Store
}

// Create ...
func (r *OptionsSoftsRepository) Create(ctx context.Context, u *models.OptionSoft) error {
	return nil
}

// Find ...
func (r *OptionsSoftsRepository) Find(ctx context.Context, id int) (*models.OptionSoft, error) {
	return nil, nil
}

// Delete ...
func (r *OptionsSoftsRepository) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (r *OptionsSoftsRepository) Update(ctx context.Context, id int, u *models.OptionSoft) error {
	return nil
}
