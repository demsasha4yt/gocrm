package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// CustomersRepository ...
type CustomersRepository struct {
	store *Store
}

// Create ...
func (r *CustomersRepository) Create(ctx context.Context, u *models.Customer) error {
	return nil
}

// Find ...
func (r *CustomersRepository) Find(ctx context.Context, id int) (*models.Customer, error) {
	return nil, nil
}

// Delete ...
func (r *CustomersRepository) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (r *CustomersRepository) Update(ctx context.Context, id int, u *models.Customer) error {
	return nil
}
