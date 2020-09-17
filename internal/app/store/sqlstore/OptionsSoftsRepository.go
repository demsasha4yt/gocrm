package sqlstore

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// OptionsSoftsRepository ...
type OptionsSoftsRepository struct {
	store *Store
}

// Create ...
func (r *OptionsSoftsRepository) Create(u *models.OptionSoft) error {
	return nil
}

// Find ...
func (r *OptionsSoftsRepository) Find(id int) (*models.OptionSoft, error) {
	return nil, nil
}

// Delete ...
func (r *OptionsSoftsRepository) Delete(id int) error {
	return nil
}

// Update ...
func (r *OptionsSoftsRepository) Update(id int, u *models.OptionSoft) error {
	return nil
}
