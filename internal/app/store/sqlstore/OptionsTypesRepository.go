package sqlstore

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// OptionsTypesRepository struct { ...
type OptionsTypesRepository struct {
	store *Store
}

// Create ...
func (r *OptionsTypesRepository) Create(u *models.OptionType) error {
	return nil
}

// Find ...
func (r *OptionsTypesRepository) Find(id int) (*models.OptionType, error) {
	return nil, nil
}

// Delete ...
func (r *OptionsTypesRepository) Delete(id int) error {
	return nil
}

// Update ...
func (r *OptionsTypesRepository) Update(id int, u *models.OptionType) error {
	return nil
}
