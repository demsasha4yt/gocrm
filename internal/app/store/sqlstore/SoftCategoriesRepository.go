package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// SoftCategoriesRepository ...
type SoftCategoriesRepository struct {
	store *Store
}

// Create ...
func (r *SoftCategoriesRepository) Create(ctx context.Context, u *models.SoftCategory) error {
	return r.store.db.QueryRow(
		ctx,
		"INSERT INTO softs_categories (name, value) VALUES ($1, $2) RETURNING id",
		u.Name,
		u.Value,
	).Scan(
		&u.ID,
	)
}

// Find ...
func (r *SoftCategoriesRepository) Find(ctx context.Context, id int) (*models.SoftCategory, error) {
	u := &models.SoftCategory{}

	if err := r.store.db.QueryRow(
		ctx,
		`SELECT id, name, value FROM softs_categories WHERE id = $1`,
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Value,
	); err != nil {
		return nil, err
	}
	return u, nil
}

// FindAll ...
func (r *SoftCategoriesRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.SoftCategory, error) {
	var m []*models.SoftCategory = make([]*models.SoftCategory, 0)
	rows, err := r.store.db.Query(
		ctx,
		`SELECT id, name, value FROM softs_categories OFFSET $1 LIMIT $2`,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := &models.SoftCategory{}
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Value,
		); err != nil {
			return nil, err
		}
		m = append(m, u)
	}
	return m, nil
}

// Delete ...
func (r *SoftCategoriesRepository) Delete(ctx context.Context, id int) error {
	_, err := r.store.db.Exec(ctx, "DELETE FROM softs_categories WHERE id=$1 ", id)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *SoftCategoriesRepository) Update(ctx context.Context, id int, u *models.SoftCategory) error {
	softCategoriesDetails, err := r.Find(ctx, id)
	if err != nil {
		return err
	}
	if u.Name != "" {
		softCategoriesDetails.Name = u.Name
	}
	if u.Value != 0 {
		softCategoriesDetails.Value = u.Value
	}
	_, err = r.store.db.Exec(
		ctx,
		`UPDATE softs_categories SET(name, value) = ($1, $2) WHERE id = $3`,
		softCategoriesDetails.Name,
		softCategoriesDetails.Value,
		id,
	)
	if err != nil {
		return err
	}
	u = softCategoriesDetails
	return nil
}
