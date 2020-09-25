package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// OptionsTypesRepository struct { ...
type OptionsTypesRepository struct {
	store *Store
}

// Create ...
func (r *OptionsTypesRepository) Create(ctx context.Context, u *models.OptionType) error {
	return r.store.db.QueryRow(
		ctx,
		"INSERT INTO options_types (name, is_soft) VALUES ($1, $2) RETURNING id",
		u.Name,
		u.IsSoft,
	).Scan(
		&u.ID,
	)
}

// Find ...
func (r *OptionsTypesRepository) Find(ctx context.Context, id int) (*models.OptionType, error) {
	u := &models.OptionType{}
	if err := r.store.db.QueryRow(
		ctx,
		`SELECT id, name, is_soft FROM options_types WHERE id = $1 LIMIT 1`,
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.IsSoft,
	); err != nil {
		return nil, err
	}
	return u, nil
}

// FindAll ...
func (r *OptionsTypesRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.OptionType, error) {
	var m []*models.OptionType = make([]*models.OptionType, 0)
	rows, err := r.store.db.Query(
		ctx,
		`SELECT id, name, is_soft FROM options_types OFFSET $1 LIMIT $2`,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := &models.OptionType{}
		if err := rows.Scan(&u.ID,
			&u.Name,
			&u.IsSoft,
		); err != nil {
			return nil, err
		}
		m = append(m, u)
	}
	return m, nil
}

// Delete ...
func (r *OptionsTypesRepository) Delete(ctx context.Context, id int) error {
	_, err := r.store.db.Exec(ctx, "DELETE FROM options_types WHERE id=$1 ", id)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *OptionsTypesRepository) Update(ctx context.Context, id int, u *models.OptionType) error {
	optionTypeDetails, err := r.Find(ctx, id)
	if err != nil {
		return err
	}
	if u.Name != "" {
		optionTypeDetails.Name = u.Name
	}
	if u.IsSoft.Ptr() != nil {
		optionTypeDetails.IsSoft = u.IsSoft
	}
	_, err = r.store.db.Exec(
		ctx,
		`UPDATE options_types SET(name, is_soft) = ($1, $2) WHERE id = $3`,
		optionTypeDetails.Name,
		optionTypeDetails.IsSoft,
		id,
	)
	if err != nil {
		return err
	}
	u = optionTypeDetails
	return nil
}
