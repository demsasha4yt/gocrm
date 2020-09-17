package sqlstore

import (
	"context"
	"database/sql"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
)

// UnitRepository is a repository for units
type UnitRepository struct {
	store *Store
}

// Create ..
func (r *UnitRepository) Create(u *models.Unit) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return r.store.db.QueryRow(context.Background(), "INSERT INTO units(name, address) VALUES($1, $2) RETURNING id",
		u.Name,
		u.Address,
	).Scan(&u.ID)
}

// Find ...
func (r *UnitRepository) Find(id int) (*models.Unit, error) {
	u := &models.Unit{}
	if err := r.store.db.QueryRow(
		context.Background(),
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
func (r *UnitRepository) Delete(id int) error {
	_, err := r.store.db.Exec(context.Background(), "DELETE FROM units WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *UnitRepository) Update(id int, u *models.Unit) error {
	if err := u.Validate(); err != nil {
		return err
	}
	unitDetails, err := r.Find(id)
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
		context.Background(),
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
