package sqlstore

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// ManufacturersRepository ...
type ManufacturersRepository struct {
	store *Store
}

// Create ...
func (r *ManufacturersRepository) Create(u *models.Manufacturer) error {
	return nil
}

// Find ...
func (r *ManufacturersRepository) Find(id int) (*models.Manufacturer, error) {
	return nil, nil
}

// Delete ...
func (r *ManufacturersRepository) Delete(id int) error {
	return nil
}

// Update ...
func (r *ManufacturersRepository) Update(id int, u *models.Manufacturer) error {
	return nil
}
