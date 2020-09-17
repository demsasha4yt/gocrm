package sqlstore

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// OrdersRepository ...
type OrdersRepository struct {
	store *Store
}

// Create ...
func (r *OrdersRepository) Create(u *models.Order) error {
	return nil
}

// Find ...
func (r *OrdersRepository) Find(id int) (*models.Order, error) {
	return nil, nil
}

// Delete ...
func (r *OrdersRepository) Delete(id int) error {
	return nil
}

// Update ...
func (r *OrdersRepository) Update(id int, u *models.Order) error {
	return nil
}
