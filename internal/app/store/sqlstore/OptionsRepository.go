package sqlstore

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// OptionsRepository ...
type OptionsRepository struct {
	store *Store
}

// Create ...
func (r *OptionsRepository) Create(u *models.Option) error {
	return nil
}

// Find ...
func (r *OptionsRepository) Find(id int) (*models.Option, error) {
	return nil, nil
}

// Delete ...
func (r *OptionsRepository) Delete(id int) error {
	return nil
}

// Update ...
func (r *OptionsRepository) Update(id int, u *models.Option) error {
	return nil
}
