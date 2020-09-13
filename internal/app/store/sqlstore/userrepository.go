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

func findUnits(r *UserRepository, u *models.User) error {
	units, err := r.store.Unit().FindUnitsByUserID(u.ID)
	if err != nil {
		return err
	}
	u.Units = units
	return nil
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
		"SELECT id, login, password, email, first_name, last_name, third_name, access_level FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.ThirdName,
		&u.AccessLevel,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	if err := findUnits(r, u); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail user by Email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, password, email, first_name, last_name, third_name, access_level FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.ThirdName,
		&u.AccessLevel,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	if err := findUnits(r, u); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByLogin user by Login
func (r *UserRepository) FindByLogin(login string) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, password, email, first_name, last_name, third_name, access_level FROM users WHERE login = $1",
		login,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.ThirdName,
		&u.AccessLevel,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	if err := findUnits(r, u); err != nil {
		return nil, err
	}
	return u, nil
}

// Delete deletes user
func (r *UserRepository) Delete(id int) error {
	_, err := r.store.db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Update user
func (r *UserRepository) Update(id int, u *models.User) error {
	u.EncryptedPassword = "a" // to avoid validate passoword field
	if err := u.Validate(); err != nil {
		return err
	}

	userDetails, err := r.Find(id)
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
	if u.AccessLevel != 0 {
		userDetails.AccessLevel = u.AccessLevel
	}

	_, err = r.store.db.Exec(
		"UPDATE users SET(login, password, email, first_name, last_name, third_name, access_level) = ($1, $2, $3, $4, $5, $6, $7) WHERE id=$8",
		userDetails.Login,
		userDetails.EncryptedPassword,
		userDetails.Email,
		userDetails.FirstName,
		userDetails.LastName,
		userDetails.ThirdName,
		userDetails.AccessLevel,
		userDetails.ID,
	)
	if err != nil {
		return err
	}
	u = userDetails
	return nil
}
