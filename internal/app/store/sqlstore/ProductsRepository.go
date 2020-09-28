package sqlstore

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// ProductsRepository ...
type ProductsRepository struct {
	store *Store
}

// Create ...
func (r *ProductsRepository) Create(ctx context.Context, u *models.Product) error {
	tx, err := r.store.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err := r.store.db.QueryRow(
		ctx,
		"INSERT INTO products (name, manufacturer_id) VALUES ($1, $2) RETURNING id",
		u.Name,
		u.ManufacturerID,
	).Scan(
		&u.ID,
	); err != nil {
		return err
	}
	for _, category := range u.Categories {
		if _, err := tx.Exec(
			ctx,
			"INSERT INTO products_categories (product_id, category_id) VALUES ($1, $2)",
			u.ID, category.ID,
		); err != nil {
			return err
		}
	}
	variationsIDs := make([]int, 0)
	for _, variation := range u.Variations {
		var variationID int
		err := tx.QueryRow(
			ctx,
			"INSERT INTO variations (name, description, price, product_id) VALUES ($1, $2, $3, $4) RETURNING id",
			variation.Name,
			variation.Description,
			variation.Price,
			variation.ProductID,
		).Scan(
			&variationID,
		)
		if err != nil {
			return err
		}
		if _, err := tx.Exec(
			ctx,
			"INSERT INTO products_variations (product_id, variation_id) VALUES ($1, $2)",
			u.ID, variationID,
		); err != nil {
			return err
		}
		variationsIDs = append(variationsIDs, variationID)
	}
	for _, value := range u.OptionsValues {
		if _, err := tx.Exec(
			ctx,
			"INSERT INTO product_options_values (product_id, options_value_id) VALUES ($1, $2)",
			u.ID, value.ID,
		); err != nil {
			return err
		}
		for _, variationID := range variationsIDs {
			if _, err := tx.Exec(
				ctx,
				"INSERT INTO variations_options_values (variation_id, options_value_id) VALUES ($1, $2)",
				variationID, value.ID,
			); err != nil {
				return err
			}
		}
	}
	return tx.Commit(ctx)
}

// Find ...
func (r *ProductsRepository) Find(ctx context.Context, id int) (*models.Product, error) {
	u := &models.Product{}
	var manufacturer []byte
	var categories []byte
	var optionsValues []byte
	if err := r.store.db.QueryRow(
		ctx,
		`SELECT p.id, p.name, p.manufacturer_id, to_jsonb(m) as manufacturer,
			COALESCE(jsonb_agg(c) FILTER (WHERE c.id IS NOT NULL), '[]') AS categories,
			COALESCE(jsonb_agg(ov) FILTER (WHERE ov.id IS NOT NULL), '[]') AS options_values
		FROM products p
		LEFT JOIN manufacturers m ON m.id = p.manufacturer_id
		LEFT JOIN products_categories pc ON pc.product_id = p.id
		LEFT JOIN categories c ON c.id = pc.category_id
		LEFT JOIN products_options_values pov ON pov.product_id = p.id
		LEFT JOIN (
			SELECT ov.id, ov.value, ov.image, ov.option_id, ov.options_type_id,
			to_jsonb(ot)as options_type, to_jsonb(o) as option
			FROM options_values ov
			LEFT JOIN options_types ot ON ot.id = ov.options_type_id
			LEFT JOIN options o ON o.id = ov.option_id
			WHERE ov.id = pov.options_value_id
		) ov ON ov.id = pov.options_value_id
		WHERE p.id = $1
		GROUP BY p.id, m.id`,
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.ManufacturerID,
		&manufacturer,
		&categories,
		&optionsValues,
	); err != nil {
		return nil, err
	}
	manufacturerEntity, err := models.NewManufacturerFromByte(manufacturer)
	if err == nil && manufacturerEntity != nil {
		u.Manufacturer = manufacturerEntity
	}
	categoriesEntities, err := models.NewCategorySliceFromByte(manufacturer)
	if err == nil && categoriesEntities != nil {
		u.Categories = categoriesEntities
	}
	optionsValuesEntities, err := models.NewOptionValueSliceFromByte(optionsValues)
	if err == nil && optionsValuesEntities != nil {
		u.OptionsValues = optionsValuesEntities
	}
	return u, nil
}

// FindAll ...
func (r *ProductsRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.Product, error) {
	var m []*models.Product = make([]*models.Product, 0)
	rows, err := r.store.db.Query(
		ctx,
		`SELECT p.id, p.name, p.manufacturer_id, to_jsonb(m) as manufacturer,
		COALESCE(jsonb_agg(c) FILTER (WHERE c.id IS NOT NULL), '[]') AS categories,
		COALESCE(jsonb_agg(ov) FILTER (WHERE ov.id IS NOT NULL), '[]') AS options_values
		FROM products p
		LEFT JOIN manufacturers m ON m.id = p.manufacturer_id
		LEFT JOIN products_categories pc ON pc.product_id = p.id
		LEFT JOIN categories c ON c.id = pc.category_id
		LEFT JOIN products_options_values pov ON pov.product_id = p.id
		LEFT JOIN (
			SELECT ov.id, ov.value, ov.image, ov.option_id, ov.options_type_id,
			to_jsonb(ot)as options_type, to_jsonb(o) as option
			FROM options_values ov
			LEFT JOIN options_types ot ON ot.id = ov.options_type_id
			LEFT JOIN options o ON o.id = ov.option_id
			WHERE ov.id = pov.options_value_id
		) ov ON ov.id = pov.options_value_id
		GROUP BY p.id, m.id
		OFFSET $1 LIMIT $2`,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var manufacturer []byte
		var categories []byte
		var optionsValues []byte
		u := &models.Product{}
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.ManufacturerID,
			&manufacturer,
			&categories,
			&optionsValues,
		); err != nil {
			return nil, err
		}
		manufacturerEntity, err := models.NewManufacturerFromByte(manufacturer)
		if err == nil && manufacturerEntity != nil {
			u.Manufacturer = manufacturerEntity
		}
		categoriesEntities, err := models.NewCategorySliceFromByte(manufacturer)
		if err == nil && categoriesEntities != nil {
			u.Categories = categoriesEntities
		}
		optionsValuesEntities, err := models.NewOptionValueSliceFromByte(optionsValues)
		if err == nil && optionsValuesEntities != nil {
			u.OptionsValues = optionsValuesEntities
		}
		m = append(m, u)
	}
	return m, nil
}

// Delete ...
func (r *ProductsRepository) Delete(ctx context.Context, id int) error {
	_, err := r.store.db.Exec(ctx, "DELETE FROM products WHERE id=$1 ", id)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *ProductsRepository) Update(ctx context.Context, id int, u *models.Product) error {
	productDetails, err := r.Find(ctx, id)
	if err != nil {
		return err
	}
	if u.Name != "" {
		productDetails.Name = u.Name
	}
	if u.ManufacturerID != 0 {
		productDetails.ManufacturerID = u.ManufacturerID
	}
	_, err = r.store.db.Exec(
		ctx,
		`UPDATE products SET(name, manufacturer_id) = ($1, $2) WHERE id = $3`,
		productDetails.Name,
		productDetails.ManufacturerID,
		id,
	)
	if err != nil {
		return err
	}
	u = productDetails
	return nil
}
