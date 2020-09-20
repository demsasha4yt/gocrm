package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OrdersRepository ...
type OrdersRepository struct {
	store *Store
}

// Create ...
func (r *OrdersRepository) Create(ctx context.Context, u *models.Order) error {
	return nil
}

// Find ...
func (r *OrdersRepository) Find(ctx context.Context, id int) (*models.Order, error) {
	return nil, nil
}

// Delete ...
func (r *OrdersRepository) Delete(ctx context.Context, id int) error {
	return nil
}

// Update ...
func (r *OrdersRepository) Update(ctx context.Context, id int, u *models.Order) error {
	return nil
}
