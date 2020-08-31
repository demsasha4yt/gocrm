package sqlstore

import (
	"database/sql"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
)

// UserRepository struct
type UserRepository struct {
	store *Store
}

// Create creates user
func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, login, password, access_level) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Email,
		u.Login,
		u.EncryptedPassword,
		u.AccessLevel,
	).Scan(&u.ID)
}

// Find user
func (r *UserRepository) Find(id int) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// FindByEmail user by Email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

// FindByLogin user by Login
func (r *UserRepository) FindByLogin(login string) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, password FROM users WHERE login = $1",
		login,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
