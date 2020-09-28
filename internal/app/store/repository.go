package store

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// UsersRepository interface
type UsersRepository interface {
	Create(context.Context, *models.User) error
	Find(context.Context, int) (*models.User, error)
	FindAll(context.Context, int, int) ([]*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByLogin(context.Context, string) (*models.User, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.User) error
}

// UnitsRepository interface
type UnitsRepository interface {
	Create(context.Context, *models.Unit) error
	Find(context.Context, int) (*models.Unit, error)
	FindAll(context.Context, int, int) ([]*models.Unit, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Unit) error
}

// CategoriesRepository interface
type CategoriesRepository interface {
	Create(context.Context, *models.Category) error
	FindAll(context.Context, int, int) ([]*models.Category, error)
	Find(context.Context, int) (*models.Category, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Category) error
}

// CustomersRepository interface
type CustomersRepository interface {
	Create(context.Context, *models.Customer) error
	Find(context.Context, int) (*models.Customer, error)
	FindAll(context.Context, int, int) ([]*models.Customer, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Customer) error
}

// ManufacturersRepository interface
type ManufacturersRepository interface {
	Create(context.Context, *models.Manufacturer) error
	FindAll(context.Context, int, int) ([]*models.Manufacturer, error)
	Find(context.Context, int) (*models.Manufacturer, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Manufacturer) error
}

// OptionsRepository interface
type OptionsRepository interface {
	Create(context.Context, *models.Option) error
	Find(context.Context, int) (*models.Option, error)
	FindAll(context.Context, int, int) ([]*models.Option, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Option) error
}

// OptionsSoftsRepository interface
type OptionsSoftsRepository interface {
	Create(context.Context, *models.OptionSoft) error
	Find(context.Context, int) (*models.OptionSoft, error)
	FindAll(context.Context, int, int) ([]*models.OptionSoft, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.OptionSoft) error
}

// OptionsTypesRepository interface
type OptionsTypesRepository interface {
	Create(context.Context, *models.OptionType) error
	Find(context.Context, int) (*models.OptionType, error)
	FindAll(context.Context, int, int) ([]*models.OptionType, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.OptionType) error
}

// OptionsValuesRepository interface
type OptionsValuesRepository interface {
	Create(context.Context, *models.OptionValue) error
	Find(context.Context, int) (*models.OptionValue, error)
	FindAll(context.Context, int, int) ([]*models.OptionValue, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.OptionValue) error
}

// OrdersRepository interface
type OrdersRepository interface {
	Create(context.Context, *models.Order) error
	Find(context.Context, int) (*models.Order, error)
	FindAll(context.Context, int, int) ([]*models.Order, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Order) error
}

// ProductsRepository interface
type ProductsRepository interface {
	Create(context.Context, *models.Product) error
	Find(context.Context, int) (*models.Product, error)
	FindAll(context.Context, int, int) ([]*models.Product, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Product) error
}

// SoftCategoriesRepository interface
type SoftCategoriesRepository interface {
	Create(context.Context, *models.SoftCategory) error
	Find(context.Context, int) (*models.SoftCategory, error)
	FindAll(context.Context, int, int) ([]*models.SoftCategory, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.SoftCategory) error
}

// VariationsRepository interface
type VariationsRepository interface {
	Create(context.Context, *models.Variation) error
	Find(context.Context, int) (*models.Variation, error)
	FindAll(context.Context, int, int) ([]*models.Variation, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.Variation) error
}
