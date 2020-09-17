package sqlstore

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// CategoriesRepository ...
type CategoriesRepository struct {
	store *Store
}

// Create ...
func (r *CategoriesRepository) Create(u *models.Category) error {
	return nil
}

// Find ...
func (r *CategoriesRepository) Find(id int) (*models.Category, error) {
	return nil, nil
}

// Delete ...
func (r *CategoriesRepository) Delete(id int) error {
	return nil
}

// Update ...
func (r *CategoriesRepository) Update(id int, u *models.Category) error {
	return nil
}
