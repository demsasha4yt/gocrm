package sqlstore

import (
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Store implements Store interface
type Store struct {
	db                       *pgxpool.Pool
	categoriesRepository     *CategoriesRepository
	customersRepository      *CustomersRepository
	manufacturersRepository  *ManufacturersRepository
	optionsRepository        *OptionsRepository
	optionsSoftsRepository   *OptionsSoftsRepository
	optionsTypesRepository   *OptionsTypesRepository
	optionsValuesRepository  *OptionsValuesRepository
	ordersRepository         *OrdersRepository
	productsRepository       *ProductsRepository
	softCategoriesRepository *SoftCategoriesRepository
	unitsRepository          *UnitsRepository
	usersRepository          *UsersRepository
	variationsRepository     *VariationsRepository
}

// New creates new Store
func New(db *pgxpool.Pool) store.Store {
	return &Store{
		db: db,
	}
}

// Customers returns CustomersRepository
func (s *Store) Customers() store.CustomersRepository {
	if s.customersRepository != nil {
		return s.customersRepository
	}

	s.customersRepository = &CustomersRepository{
		store: s,
	}

	return s.customersRepository
}

// Categories returns CategoriesRepository
func (s *Store) Categories() store.CategoriesRepository {
	if s.categoriesRepository != nil {
		return s.categoriesRepository
	}

	s.categoriesRepository = &CategoriesRepository{
		store: s,
	}

	return s.categoriesRepository
}

// Manufacturers returns manufacturersRepository
func (s *Store) Manufacturers() store.ManufacturersRepository {
	if s.manufacturersRepository != nil {
		return s.manufacturersRepository
	}

	s.manufacturersRepository = &ManufacturersRepository{
		store: s,
	}

	return s.manufacturersRepository
}

// Options returns optionsRepository
func (s *Store) Options() store.OptionsRepository {
	if s.optionsRepository != nil {
		return s.optionsRepository
	}

	s.optionsRepository = &OptionsRepository{
		store: s,
	}

	return s.optionsRepository
}

// OptionsSofts returns optionsSoftsRepository
func (s *Store) OptionsSofts() store.OptionsSoftsRepository {
	if s.optionsSoftsRepository != nil {
		return s.optionsSoftsRepository
	}

	s.optionsSoftsRepository = &OptionsSoftsRepository{
		store: s,
	}

	return s.optionsSoftsRepository
}

// OptionsTypes returns optionsSoftsRepository
func (s *Store) OptionsTypes() store.OptionsTypesRepository {
	if s.optionsTypesRepository != nil {
		return s.optionsTypesRepository
	}

	s.optionsTypesRepository = &OptionsTypesRepository{
		store: s,
	}

	return s.optionsTypesRepository
}

// OptionsValues returns optionsValuesRepository
func (s *Store) OptionsValues() store.OptionsValuesRepository {
	if s.optionsValuesRepository != nil {
		return s.optionsValuesRepository
	}

	s.optionsValuesRepository = &OptionsValuesRepository{
		store: s,
	}

	return s.optionsValuesRepository
}

// Orders returns ordersRepository
func (s *Store) Orders() store.OrdersRepository {
	if s.ordersRepository != nil {
		return s.ordersRepository
	}

	s.ordersRepository = &OrdersRepository{
		store: s,
	}

	return s.ordersRepository
}

// Products returns productsRepository
func (s *Store) Products() store.ProductsRepository {
	if s.productsRepository != nil {
		return s.productsRepository
	}

	s.productsRepository = &ProductsRepository{
		store: s,
	}

	return s.productsRepository
}

// SoftCategories returns softCategoriesRepository
func (s *Store) SoftCategories() store.SoftCategoriesRepository {
	if s.softCategoriesRepository != nil {
		return s.softCategoriesRepository
	}

	s.softCategoriesRepository = &SoftCategoriesRepository{
		store: s,
	}

	return s.softCategoriesRepository
}

// Units returns unitRepository
func (s *Store) Units() store.UnitsRepository {
	if s.unitsRepository != nil {
		return s.unitsRepository
	}

	s.unitsRepository = &UnitsRepository{
		store: s,
	}

	return s.unitsRepository
}

// Users returns userRepository
func (s *Store) Users() store.UsersRepository {
	if s.usersRepository != nil {
		return s.usersRepository
	}

	s.usersRepository = &UsersRepository{
		store: s,
	}

	return s.usersRepository
}

// Variations returns VariationsRepository
func (s *Store) Variations() store.VariationsRepository {
	if s.variationsRepository != nil {
		return s.variationsRepository
	}

	s.variationsRepository = &VariationsRepository{
		store: s,
	}

	return s.variationsRepository
}
