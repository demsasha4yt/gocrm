package sqlstore

import (
	"context"
	"database/sql"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
)

// UnitsRepository is a repository for units
type UnitsRepository struct {
	store *Store
}

// Create ..
func (r *UnitsRepository) Create(ctx context.Context, u *models.Unit) error {
	return r.store.db.QueryRow(ctx, "INSERT INTO units(name, address) VALUES($1, $2) RETURNING id",
		u.Name,
		u.Address,
	).Scan(&u.ID)
}

// Find ...
func (r *UnitsRepository) Find(ctx context.Context, id int) (*models.Unit, error) {
	u := &models.Unit{}
	if err := r.store.db.QueryRow(
		ctx,
		"SELECT id, name, address FROM units WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Address,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

// Delete ..
func (r *UnitsRepository) Delete(ctx context.Context, id int) error {
	_, err := r.store.db.Exec(ctx, "DELETE FROM units WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *UnitsRepository) Update(ctx context.Context, id int, u *models.Unit) error {
	unitDetails, err := r.Find(ctx, id)
	if err != nil {
		return store.ErrRecordNotFound
	}
	if u.Name != "" {
		unitDetails.Name = u.Name
	}
	if u.Address != "" {
		unitDetails.Address = u.Address
	}
	_, err = r.store.db.Exec(
		ctx,
		"UPDATE units SET(name, address) = ($1, $2) WHERE id=$3",
		&unitDetails.Name,
		&unitDetails.Address,
		&unitDetails.ID,
	)
	if err != nil {
		return err
	}
	u = unitDetails
	return nil
}
