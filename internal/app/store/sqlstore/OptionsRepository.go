package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OptionsRepository ...
type OptionsRepository struct {
	store *Store
}

// Create ...
func (r *OptionsRepository) Create(ctx context.Context, u *models.Option) error {
	return r.store.db.QueryRow(
		ctx,
		"INSERT INTO options (name, description, options_type_id) VALUES ($1, $2, $3) RETURNING id",
		u.Name,
		u.Description,
		u.OptionTypeID,
	).Scan(
		&u.ID,
	)
}

// Find ...
func (r *OptionsRepository) Find(ctx context.Context, id int) (*models.Option, error) {
	u := &models.Option{}
	if err := r.store.db.QueryRow(
		ctx,
		`SELECT id, name, description, options_type_id FROM options WHERE id = $1`,
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Description,
		&u.OptionTypeID,
	); err != nil {
		return nil, err
	}
	return u, nil
}

// FindAll ...
func (r *OptionsRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.Option, error) {
	var m []*models.Option = make([]*models.Option, 0)
	rows, err := r.store.db.Query(
		ctx,
		`SELECT id, name, description, options_type_id FROM options OFFSET $1 LIMIT $2`,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := &models.Option{}
		if err := rows.Scan(&u.ID,
			&u.Name,
			&u.Description,
			&u.OptionTypeID,
		); err != nil {
			return nil, err
		}
		m = append(m, u)
	}
	return m, nil
}

// Delete ...
func (r *OptionsRepository) Delete(ctx context.Context, id int) error {
	_, err := r.store.db.Exec(ctx, "DELETE FROM options WHERE id=$1 ", id)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *OptionsRepository) Update(ctx context.Context, id int, u *models.Option) error {
	optionDetails, err := r.Find(ctx, id)
	if err != nil {
		return err
	}
	if u.Name != "" {
		optionDetails.Name = u.Name
	}
	if u.Description != "" {
		optionDetails.Description = u.Description
	}
	if u.OptionTypeID != 0 {
		optionDetails.OptionTypeID = u.OptionTypeID
	}
	_, err = r.store.db.Exec(
		ctx,
		`UPDATE options SET(name, description, options_type_id) = ($1, $2, $3) WHERE id = $4`,
		optionDetails.Name,
		optionDetails.Description,
		optionDetails.OptionTypeID,
		id,
	)
	if err != nil {
		return err
	}
	u = optionDetails
	return nil
}
