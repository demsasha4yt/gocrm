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
}

// CategoriesRepository interface
type CategoriesRepository interface {
	Create(*models.Category) error
	FindAll() ([]*models.Category, error)
	Find(int) (*models.Category, error)
	Delete(int) error
	Update(int, *models.Category) error
}

// CustomersRepository interface
type CustomersRepository interface {
	Create(*models.Customer) error
	Find(int) (*models.Customer, error)
	Delete(int) error
	Update(int, *models.Customer) error
}

// ManufacturersRepository interface
type ManufacturersRepository interface {
	Create(*models.Manufacturer) error
	FindAll() ([]*models.Manufacturer, error)
	Find(int) (*models.Manufacturer, error)
	Delete(int) error
	Update(int, *models.Manufacturer) error
}

// OptionsRepository interface
type OptionsRepository interface {
	Create(*models.Option) error
	Find(int) (*models.Option, error)
	Delete(int) error
	Update(int, *models.Option) error
}

// OptionsSoftsRepository interface
type OptionsSoftsRepository interface {
	Create(*models.OptionSoft) error
	Find(int) (*models.OptionSoft, error)
	Delete(int) error
	Update(int, *models.OptionSoft) error
}

// OptionsTypesRepository interface
type OptionsTypesRepository interface {
	Create(*models.OptionType) error
	Find(int) (*models.OptionType, error)
	Delete(int) error
	Update(int, *models.OptionType) error
}

// OptionsValuesRepository interface
type OptionsValuesRepository interface {
	Create(*models.OptionValue) error
	Find(int) (*models.OptionValue, error)
	Delete(int) error
	Update(int, *models.OptionValue) error
}

// OrdersRepository interface
type OrdersRepository interface {
	Create(*models.Order) error
	Find(int) (*models.Order, error)
	Delete(int) error
	Update(int, *models.Order) error
}

// ProductsRepository interface
type ProductsRepository interface {
	Create(*models.Product) error
	Find(int) (*models.Product, error)
	Delete(int) error
	Update(int, *models.Product) error
}

// VariationsRepository interface
type VariationsRepository interface {
	Create(*models.Variation) error
	Find(int) (*models.Variation, error)
	Delete(int) error
	Update(int, *models.Variation) error
}
