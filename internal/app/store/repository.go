package store

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// UserRepository interface
type UserRepository interface {
	Create(*models.User) error
	Find(int) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	FindByLogin(string) (*models.User, error)
}
