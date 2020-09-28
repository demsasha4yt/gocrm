package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
)

// OptionsSoftsRepository ...
type OptionsSoftsRepository struct {
	store *Store
}

// Create ...
func (r *OptionsSoftsRepository) Create(ctx context.Context, u *models.OptionSoft) error {
	return r.store.db.QueryRow(
		ctx,
		`INSERT INTO options_softs (name, image, manufacturer_id, 
			options_value_id, soft_category_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		u.Name,
		u.Image,
		u.ManufacturerID,
		u.OptionValueID,
		u.SoftCategoryID,
	).Scan(
		&u.ID,
	)
}

// Find ...
func (r *OptionsSoftsRepository) Find(ctx context.Context, id int) (*models.OptionSoft, error) {
	u := &models.OptionSoft{}

	var optionsValue []byte
	var softCategory []byte
	var manufacturer []byte
	err := r.store.db.QueryRow(
		ctx,
		`SELECT o.id, o.name, o.image, o.manufacturer_id, o.options_value_id, o.soft_category_id,
		to_jsonb(ov) as option_value, to_jsonb(sc) as soft_category, to_jsonb(m) as manufacturer
		FROM options_softs o
		JOIN options_values ov ON ov.id = o.options_value_id
		JOIN softs_categories sc ON sc.id = o.soft_category_id
		JOIN manufacturers m ON m.id = o.manufacturer_id
		WHERE o.id = $1 LIMIT 1`,
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Image,
		&u.ManufacturerID,
		&u.OptionValueID,
		&u.SoftCategoryID,
		&optionsValue,
		&softCategory,
		&manufacturer,
	)
	if err != nil {
		return nil, err
	}
	softCategoryEntity, err := models.NewSoftCategoryFromByte(softCategory)
	if err == nil && softCategoryEntity != nil {
		u.SoftCategory = softCategoryEntity
	}
	optionValueEntity, err := models.NewOptionValueFromByte(optionsValue)
	if err == nil && optionValueEntity != nil {
		u.OptionValue = optionValueEntity
	}
	manufacturerEntity, err := models.NewManufacturerFromByte(manufacturer)
	if err == nil && manufacturerEntity != nil {
		u.Manufacturer = manufacturerEntity
	}
	return u, nil
}

// FindAll ...
func (r *OptionsSoftsRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.OptionSoft, error) {
	var m []*models.OptionSoft = make([]*models.OptionSoft, 0)
	rows, err := r.store.db.Query(
		ctx,
		`SELECT o.id, o.name, o.image, o.manufacturer_id, o.options_value_id, o.soft_category_id,
		to_jsonb(ov) as option_value, to_jsonb(sc) as soft_category, to_jsonb(m) as manufacturer
		FROM options_softs o
		JOIN options_values ov ON ov.id = o.options_value_id
		JOIN softs_categories sc ON sc.id = o.soft_category_id
		JOIN manufacturers m ON m.id = o.manufacturer_id
		OFFSET $1 LIMIT $2`,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var optionsValue []byte
		var softCategory []byte
		var manufacturer []byte
		u := &models.OptionSoft{}
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Image,
			&u.ManufacturerID,
			&u.OptionValueID,
			&u.SoftCategoryID,
			&optionsValue,
			&softCategory,
			&manufacturer,
		); err != nil {
			return nil, err
		}
		softCategoryEntity, err := models.NewSoftCategoryFromByte(softCategory)
		if err == nil && softCategoryEntity != nil {
			u.SoftCategory = softCategoryEntity
		}
		optionValueEntity, err := models.NewOptionValueFromByte(optionsValue)
		if err == nil && optionValueEntity != nil {
			u.OptionValue = optionValueEntity
		}
		manufacturerEntity, err := models.NewManufacturerFromByte(manufacturer)
		if err == nil && manufacturerEntity != nil {
			u.Manufacturer = manufacturerEntity
		}
		m = append(m, u)
	}
	return m, nil
}

// Delete ...
func (r *OptionsSoftsRepository) Delete(ctx context.Context, id int) error {
	_, err := r.store.db.Exec(ctx, "DELETE FROM options_softs WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *OptionsSoftsRepository) Update(ctx context.Context, id int, u *models.OptionSoft) error {
	optionSoftDetails, err := r.Find(ctx, id)
	if err != nil {
		return store.ErrRecordNotFound
	}
	if u.Name != "" {
		optionSoftDetails.Name = u.Name
	}
	if u.Image != "" {
		optionSoftDetails.Image = u.Image
	}
	if u.ManufacturerID != 0 {
		optionSoftDetails.ManufacturerID = u.ManufacturerID
	}
	if u.OptionValueID != 0 {
		optionSoftDetails.OptionValueID = u.OptionValueID
	}
	if u.SoftCategoryID != 0 {
		optionSoftDetails.SoftCategoryID = u.SoftCategoryID
	}
	_, err = r.store.db.Exec(
		ctx,
		`UPDATE options_softs SET(name, image, manufacturer_id, 
			options_value_id, soft_category_id) = ($1, $2, $3, $4,$5) WHERE id=$6`,
		&optionSoftDetails.Name,
		&optionSoftDetails.Image,
		&optionSoftDetails.ManufacturerID,
		&optionSoftDetails.OptionValueID,
		&optionSoftDetails.SoftCategoryID,
		&optionSoftDetails.ID,
	)
	if err != nil {
		return err
	}
	u = optionSoftDetails
	return nil
}
