package sqlstore

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// VariationsRepository ...
type VariationsRepository struct {
	store *Store
}

// Create ...
func (r *VariationsRepository) Create(u *models.Variation) error {
	return nil
}

// Find ...
func (r *VariationsRepository) Find(id int) (*models.Variation, error) {
	return nil, nil
}

// Delete ...
func (r *VariationsRepository) Delete(id int) error {
	return nil
}

// Update ...
func (r *VariationsRepository) Update(id int, u *models.Variation) error {
	return nil
}
