package sqlstore

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// ProductsRepository ...
type ProductsRepository struct {
	store *Store
}

// Create ...
func (r *ProductsRepository) Create(u *models.Product) error {
	return nil
}

// Find ...
func (r *ProductsRepository) Find(id int) (*models.Product, error) {
	return nil, nil
}

// Delete ...
func (r *ProductsRepository) Delete(id int) error {
	return nil
}

// Update ...
func (r *ProductsRepository) Update(id int, u *models.Product) error {
	return nil
}
