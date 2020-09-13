package store

import "github.com/demsasha4yt/gocrm.git/internal/app/models"

// UserRepository interface
type UserRepository interface {
	Create(*models.User) error
	Find(int) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	FindByLogin(string) (*models.User, error)
	Delete(int) error
	Update(int, *models.User) error
}

// UnitRepository interface
type UnitRepository interface {
	Create(*models.Unit) error
	Find(int) (*models.Unit, error)
	Delete(int) error
	Update(int, *models.Unit) error
	FindUnitsByUserID(int) ([]*models.Unit, error)
}
