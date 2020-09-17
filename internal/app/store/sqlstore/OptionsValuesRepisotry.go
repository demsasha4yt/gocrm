package sqlstore

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// OptionsValuesRepository ...
type OptionsValuesRepository struct {
	store *Store
}

// Create ...
func (r *OptionsValuesRepository) Create(u *models.OptionValue) error {
	return nil
}

// Find ...
func (r *OptionsValuesRepository) Find(id int) (*models.OptionValue, error) {
	return nil, nil
}

// Delete ...
func (r *OptionsValuesRepository) Delete(id int) error {
	return nil
}

// Update ...
func (r *OptionsValuesRepository) Update(id int, u *models.OptionValue) error {
	return nil
}
