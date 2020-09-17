package sqlstore

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
)

// ManufacturersRepository ...
type ManufacturersRepository struct {
	store *Store
}

// Create ...
func (r *ManufacturersRepository) Create(u *models.Manufacturer) error {
	if err := u.Validate(); err != nil {
		return err
	}
	tx, err := r.store.db.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	if err := tx.QueryRow(
		context.Background(),
		"INSERT INTO manufacturers (name, description) VALUES ($1, $2) RETURNING id",
		u.Name,
		u.Description,
	).Scan(
		&u.ID,
	); err != nil {
		return err
	}

	for _, unit := range u.Units {
		if _, err := tx.Exec(
			context.Background(),
			"INSERT INTO manufacturers_units(manufacturer_id, unit_id) VALUES ($1, $2)",
			u.ID, unit.ID,
		); err != nil {
			return err
		}
	}

	return tx.Commit(context.Background())
}

// FindAll ...
func (r *ManufacturersRepository) FindAll(page int) ([]*models.Manufacturer, error) {
	var m []*models.Manufacturer = make([]*models.Manufacturer, 0)

	offset := (page - 1) * 50
	limit := 50

	rows, err := r.store.db.Query(
		context.Background(),
		`SELECT id, name, description FROM manufacturers OFFSET $1 LIMIT $2`,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := &models.Manufacturer{}
		if err := rows.Scan(&u.ID,
			&u.Name,
			&u.Description,
		); err != nil {
			return nil, err
		}
		m = append(m, u)
	}
	return m, nil
}

// Find ...
func (r *ManufacturersRepository) Find(id int) (*models.Manufacturer, error) {
	u := &models.Manufacturer{}
	var units []byte

	if err := r.store.db.QueryRow(
		context.Background(),
		`SELECT m.id, m.name, m.description,
			COALESCE(json_agg(u) FILTER (WHERE u.id IS NOT NULL), '[]') AS units
		FROM manufacturers m
		LEFT JOIN manufacturers_units mu ON mu.manufacturer_id = m.id
		LEFT JOIN units u on u.id = mu.unit_id
		WHERE m.id = $1
		GROUP BY m.id`,
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Description,
		&units,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	if err := json.Unmarshal(units, &u.Units); err != nil {
		return nil, err
	}
	return u, nil
}

// Delete ...
func (r *ManufacturersRepository) Delete(id int) error {
	_, err := r.store.db.Exec(context.Background(), "DELETE FROM manufacturers WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *ManufacturersRepository) Update(id int, u *models.Manufacturer) error {
	if err := u.Validate(); err != nil {
		return err
	}
	manufacturerDetails, err := r.Find(id)
	if err != nil {
		return store.ErrRecordNotFound
	}
	if u.Name != "" {
		manufacturerDetails.Name = u.Name
	}
	if u.Description != "" {
		manufacturerDetails.Description = u.Description
	}
	_, err = r.store.db.Exec(
		context.Background(),
		"UPDATE manufacturers SET(name, description) = ($1, $2) WHERE id=$3",
		&manufacturerDetails.Name,
		&manufacturerDetails.Description,
		&manufacturerDetails.ID,
	)
	if err != nil {
		return err
	}
	u = manufacturerDetails
	return nil
}
