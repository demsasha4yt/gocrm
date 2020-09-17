package sqlstore

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// CustomersRepository ...
type CustomersRepository struct {
	store *Store
}

// Create ...
func (r *CustomersRepository) Create(u *models.Customer) error {
	return nil
}

// Find ...
func (r *CustomersRepository) Find(id int) (*models.Customer, error) {
	return nil, nil
}

// Delete ...
func (r *CustomersRepository) Delete(id int) error {
	return nil
}

// Update ...
func (r *CustomersRepository) Update(id int, u *models.Customer) error {
	return nil
}
