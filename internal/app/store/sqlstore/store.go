package sqlstore

import (
	"database/sql"

	"github.com/demsasha4yt/gocrm.git/internal/app/store"
)

// Store implements Store interface
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New creates new Store
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User returns userRepository
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
