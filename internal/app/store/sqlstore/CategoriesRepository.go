package sqlstore

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
)

// CategoriesRepository ...
type CategoriesRepository struct {
	store *Store
}

// Create ...
func (r *CategoriesRepository) Create(ctx context.Context, u *models.Category) error {
	return r.store.db.QueryRow(
		ctx,
		"INSERT INTO categories (name, description, parent_id) VALUES ($1, $2, $3) RETURNING id",
		u.Name,
		u.Description,
		u.ParentID,
	).Scan(&u.ID)
}

// FindAll ...
func (r *CategoriesRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.Category, error) {
	var categories []*models.Category = make([]*models.Category, 0)

	rows, err := r.store.db.Query(
		ctx,
		`SELECT c.id, c.name, c.description, c.parent_id,
		COALESCE(jsonb_agg(cc) FILTER (WHERE cc.id IS NOT NULL), '[]') AS subcategories
		FROM categories c
		LEFT JOIN categories cc ON cc.parent_id = c.id
		GROUP BY c.id
		OFFSET $1 LIMIT $2`,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var subs []byte
		u := &models.Category{}
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Description,
			&u.ParentID,
			&subs,
		); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(subs, &u.Subcategories); err != nil {
			return nil, err
		}
		categories = append(categories, u)
	}
	return categories, nil
}

// Find ...
func (r *CategoriesRepository) Find(ctx context.Context, id int) (*models.Category, error) {
	u := &models.Category{}
	var subs []byte

	if err := r.store.db.QueryRow(
		ctx,
		`SELECT c.id, c.name, c.description, c.parent_id,
		COALESCE(jsonb_agg(cc) FILTER (WHERE cc.id IS NOT NULL), '[]') AS subcategories
		FROM categories c
		LEFT JOIN categories cc ON cc.parent_id = c.id
		WHERE c.id = $1
		GROUP BY c.id`,
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Description,
		&u.ParentID,
		&subs,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	if err := json.Unmarshal(subs, &u.Subcategories); err != nil {
		return nil, err
	}

	return u, nil
}

// Delete ...
func (r *CategoriesRepository) Delete(ctx context.Context, id int) error {
	_, err := r.store.db.Exec(ctx, "DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *CategoriesRepository) Update(ctx context.Context, id int, u *models.Category) error {
	categoryDetails, err := r.Find(ctx, id)
	if err != nil {
		return store.ErrRecordNotFound
	}
	if u.Name != "" {
		categoryDetails.Name = u.Name
	}
	if u.Description != "" {
		categoryDetails.Description = u.Description
	}
	categoryDetails.ParentID = u.ParentID
	_, err = r.store.db.Exec(
		ctx,
		"UPDATE categories SET(name, description, parent_id) = ($1, $2, $3) WHERE id=$4",
		&categoryDetails.Name,
		&categoryDetails.Description,
		&categoryDetails.ParentID,
		&categoryDetails.ID,
	)
	if err != nil {
		return err
	}
	u = categoryDetails
	return nil
}
