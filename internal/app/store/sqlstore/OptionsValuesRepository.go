package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
)

// OptionsValuesRepository ...
type OptionsValuesRepository struct {
	store *Store
}

// Create ...
func (r *OptionsValuesRepository) Create(ctx context.Context, u *models.OptionValue) error {
	return r.store.db.QueryRow(
		ctx,
		`INSERT INTO options_values(value, image, option_id, options_type_id)
		VALUES ($1, $2, $3, $4) RETURNING id`,
		u.Value,
		u.Image,
		u.OptionID,
		u.OptionTypeID,
	).Scan(
		&u.ID,
	)
}

// Find ...
func (r *OptionsValuesRepository) Find(ctx context.Context, id int) (*models.OptionValue, error) {
	u := &models.OptionValue{}

	var option []byte
	var optionType []byte
	var softs []byte
	if err := r.store.db.QueryRow(
		ctx,
		`SELECT ov.id, ov.value, ov.image, ov.option_id, ov.options_type_id, to_jsonb(o) as option,
		to_jsonb(ot) as option_type, COALESCE(jsonb_agg(os) FILTER (WHERE os.id IS NOT NULL), '[]') AS softs
		FROM options_values ov
		JOIN options_types ot ON ot.id = ov.options_type_id
		JOIN options o ON o.id = ov.id
		LEFT JOIN options_softs os ON os.options_value_id = ov.id
		WHERE ov.id = $1
		GROUP BY ov.id, o.id, ot.id
		LIMIT 1`,
		id,
	).Scan(
		&u.ID,
		&u.Value,
		&u.Image,
		&u.OptionID,
		&u.OptionTypeID,
		&option,
		&optionType,
		&softs,
	); err != nil {
		return nil, err
	}
	optionEntity, err := models.NewOptionFromByte(option)
	if err == nil && optionEntity != nil {
		u.Option = optionEntity
	}

	optionTypeEntity, err := models.NewOptionTypeFromByte(optionType)
	if err == nil && optionTypeEntity != nil {
		u.OptionType = optionTypeEntity
	}

	softsEntities, err := models.NewOptionSoftSliceFromByte(softs)
	if err == nil && softsEntities != nil {
		u.Softs = softsEntities
	}
	return u, nil
}

// FindAll ...
func (r *OptionsValuesRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.OptionValue, error) {
	var m []*models.OptionValue = make([]*models.OptionValue, 0)

	rows, err := r.store.db.Query(
		ctx,
		`SELECT ov.id, ov.value, ov.image, ov.option_id, ov.options_type_id, to_jsonb(o) as option,
		to_jsonb(ot) as option_type, COALESCE(jsonb_agg(os) FILTER (WHERE os.id IS NOT NULL), '[]') AS softs
		FROM options_values ov
		JOIN options_types ot ON ot.id = ov.options_type_id
		JOIN options o ON o.id = ov.id
		LEFT JOIN options_softs os ON os.options_value_id = ov.id
		GROUP BY ov.id, o.id, ot.id
		OFFSET $1 LIMIT $2`,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var option []byte
		var optionType []byte
		var softs []byte
		u := &models.OptionValue{}
		if err := rows.Scan(
			&u.ID,
			&u.Value,
			&u.Image,
			&u.OptionID,
			&u.OptionTypeID,
			&option,
			&optionType,
			&softs,
		); err != nil {
			return nil, err
		}
		optionEntity, err := models.NewOptionFromByte(option)
		if err == nil && optionEntity != nil {
			u.Option = optionEntity
		}

		optionTypeEntity, err := models.NewOptionTypeFromByte(optionType)
		if err == nil && optionTypeEntity != nil {
			u.OptionType = optionTypeEntity
		}

		softsEntities, err := models.NewOptionSoftSliceFromByte(softs)
		if err == nil && softsEntities != nil {
			u.Softs = softsEntities
		}
		m = append(m, u)
	}
	return m, nil
}

// Delete ...
func (r *OptionsValuesRepository) Delete(ctx context.Context, id int) error {
	_, err := r.store.db.Exec(ctx, "DELETE FROM options_values WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *OptionsValuesRepository) Update(ctx context.Context, id int, u *models.OptionValue) error {
	optionValueDetails, err := r.Find(ctx, id)
	if err != nil {
		return store.ErrRecordNotFound
	}
	if u.Value != "" {
		optionValueDetails.Value = u.Value
	}
	if u.Image != "" {
		optionValueDetails.Image = u.Image
	}
	if u.OptionID != 0 {
		optionValueDetails.OptionID = u.OptionID
	}
	if u.OptionTypeID != 0 {
		optionValueDetails.OptionTypeID = u.OptionTypeID
	}

	_, err = r.store.db.Exec(
		ctx,
		"UPDATE options_values SET(value, image, option_id, options_type_id) = ($1, $2, $3, $4) WHERE id=$5",
		&optionValueDetails.Value,
		&optionValueDetails.Image,
		&optionValueDetails.OptionID,
		&optionValueDetails.OptionTypeID,
		&optionValueDetails.ID,
	)
	if err != nil {
		return err
	}
	u = optionValueDetails
	return nil
}
