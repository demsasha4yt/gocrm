package sqlstore

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
)

// UserRepository struct
type UserRepository struct {
	store *Store
}

// Create creates user
func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		ctx,
		"INSERT INTO users (email, login, password, access_level) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Email,
		u.Login,
		u.EncryptedPassword,
		u.AccessLevelID,
	).Scan(&u.ID)
}

// Find user
func (r *UserRepository) Find(ctx context.Context, id int) (*models.User, error) {
	u := &models.User{}

	var units []byte
	var accessLevel []byte
	if err := r.store.db.QueryRow(
		ctx,
		`SELECT users.id, users.login, users.password, users.email, users.first_name, users.last_name, users.third_name, users.access_level, 
		COALESCE(jsonb_agg(units) FILTER (WHERE units.id IS NOT NULL), '[]') AS units,
		to_jsonb(al) as access_level
		FROM users
		LEFT JOIN users_units uu ON uu.user_id = users.id
		LEFT JOIN units ON units.id = uu.unit_id
		LEFT JOIN access_levels al ON al.id = users.access_level
		WHERE users.id = $1
		GROUP BY users.id, al.id
		LIMIT 1`,
		id,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.ThirdName,
		&u.AccessLevelID,
		&units,
		&accessLevel,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	if err := json.Unmarshal(units, &u.Units); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(accessLevel, &u.AccessLevel); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail user by Email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{}

	var units []byte
	var accessLevel []byte
	if err := r.store.db.QueryRow(
		ctx,
		`SELECT users.id, users.login, users.password, users.email, users.first_name, users.last_name, users.third_name, users.access_level, 
		COALESCE(jsonb_agg(units) FILTER (WHERE units.id IS NOT NULL), '[]') AS units,
		to_jsonb(al) as access_level
		FROM users
		LEFT JOIN users_units uu ON uu.user_id = users.id
		LEFT JOIN units ON units.id = uu.unit_id
		LEFT JOIN access_levels al ON al.id = users.access_level
		WHERE users.login = $1
		GROUP BY users.id, al.id
		LIMIT 1`,
		email,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.ThirdName,
		&u.AccessLevelID,
		&units,
		&accessLevel,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	if err := json.Unmarshal(units, &u.Units); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(accessLevel, &u.AccessLevel); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByLogin user by Login
func (r *UserRepository) FindByLogin(ctx context.Context, login string) (*models.User, error) {
	u := &models.User{}

	var units []byte
	var accessLevel []byte
	if err := r.store.db.QueryRow(
		ctx,
		`SELECT users.id, users.login, users.password, users.email, users.first_name, users.last_name, users.third_name, users.access_level, 
			COALESCE(jsonb_agg(units) FILTER (WHERE units.id IS NOT NULL), '[]') AS units
		FROM users
		LEFT JOIN users_units uu ON uu.user_id = users.id
		LEFT JOIN units ON units.id = uu.unit_id
		WHERE users.login = $1
		GROUP BY users.id
		LIMIT 1`,
		login,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.ThirdName,
		&u.AccessLevelID,
		&units,
		&accessLevel,
	); err != nil {
		fmt.Printf("%+v", err)
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	if err := json.Unmarshal(units, &u.Units); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(accessLevel, &u.AccessLevel); err != nil {
		return nil, err
	}
	return u, nil
}

// Delete deletes user
func (r *UserRepository) Delete(ctx context.Context, id int) error {
	_, err := r.store.db.Exec(ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Update user
func (r *UserRepository) Update(ctx context.Context, id int, u *models.User) error {
	u.EncryptedPassword = "a" // to avoid validate passoword field
	if err := u.Validate(); err != nil {
		return err
	}

	userDetails, err := r.Find(ctx, id)
	if err != nil {
		return store.ErrRecordNotFound
	}

	if u.Login != "" {
		userDetails.Login = u.Login
	}
	if u.Email != "" {
		userDetails.Email = u.Email
	}
	if u.FirstName != "" {
		userDetails.FirstName = u.FirstName
	}
	if u.LastName != "" {
		userDetails.LastName = u.LastName
	}
	if u.ThirdName != "" {
		userDetails.ThirdName = u.ThirdName
	}
	if u.AccessLevelID != 0 {
		userDetails.AccessLevelID = u.AccessLevelID
	}

	_, err = r.store.db.Exec(
		ctx,
		"UPDATE users SET(login, password, email, first_name, last_name, third_name, access_level) = ($1, $2, $3, $4, $5, $6, $7) WHERE id=$8",
		userDetails.Login,
		userDetails.EncryptedPassword,
		userDetails.Email,
		userDetails.FirstName,
		userDetails.LastName,
		userDetails.ThirdName,
		userDetails.AccessLevelID,
		userDetails.ID,
	)
	if err != nil {
		return err
	}
	u = userDetails
	return nil
}
