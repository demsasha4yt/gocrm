package service

import "github.com/demsasha4yt/gocrm.git/internal/app/store"

// Service ...
type Service struct {
	store                store.Store
	categoriesService    *CategoriesService
	customersService     *CustomersService
	manufacturersService *ManufacturersService
	optionsService       *OptionsService
	optionsSoftsService  *OptionsSoftsService
	optionsTypesService  *OptionsTypesService
	optionsValuesService *OptionsValuesService
	ordersService        *OrdersService
	productsService      *ProductsService
	usersService         *UsersService
	unitsService         *UnitsService
	variationsService    *VariationsService
}

// New returns new service
func New(store store.Store) *Service {
	return &Service{
		store: store,
	}
}

// Customers returns CustomersService
func (s *Service) Customers() CustomersServiceInterface {
	if s.customersService != nil {
		return s.customersService
	}

	s.customersService = &CustomersService{
		service: s,
	}

	return s.customersService
}

// Categories returns CategoriesService
func (s *Service) Categories() CategoriesServiceInterface {
	if s.categoriesService != nil {
		return s.categoriesService
	}

	s.categoriesService = &CategoriesService{
		service: s,
	}

	return s.categoriesService
}

// Manufacturers returns manufacturersService
func (s *Service) Manufacturers() ManufacturersServiceInterface {
	if s.manufacturersService != nil {
		return s.manufacturersService
	}

	s.manufacturersService = &ManufacturersService{
		service: s,
	}

	return s.manufacturersService
}

// Options returns optionsService
func (s *Service) Options() OptionsServiceInterface {
	if s.optionsService != nil {
		return s.optionsService
	}

	s.optionsService = &OptionsService{
		service: s,
	}

	return s.optionsService
}

// OptionsSofts returns optionsSoftsService
func (s *Service) OptionsSofts() OptionsSoftsServiceInterface {
	if s.optionsSoftsService != nil {
		return s.optionsSoftsService
	}

	s.optionsSoftsService = &OptionsSoftsService{
		service: s,
	}

	return s.optionsSoftsService
}

// OptionsTypes returns optionsSoftsService
func (s *Service) OptionsTypes() OptionsTypesServiceInterface {
	if s.optionsTypesService != nil {
		return s.optionsTypesService
	}

	s.optionsTypesService = &OptionsTypesService{
		service: s,
	}

	return s.optionsTypesService
}

// OptionsValues returns optionsValuesService
func (s *Service) OptionsValues() OptionsValuesServiceInterface {
	if s.optionsValuesService != nil {
		return s.optionsValuesService
	}

	s.optionsValuesService = &OptionsValuesService{
		service: s,
	}

	return s.optionsValuesService
}

// Orders returns ordersService
func (s *Service) Orders() OrdersServiceInterface {
	if s.ordersService != nil {
		return s.ordersService
	}

	s.ordersService = &OrdersService{
		service: s,
	}

	return s.ordersService
}

// Products returns producrsService
func (s *Service) Products() ProductsServiceInterface {
	if s.productsService != nil {
		return s.productsService
	}

	s.productsService = &ProductsService{
		service: s,
	}

	return s.productsService
}

// Units returns unitsService
func (s *Service) Units() UnitsServiceInterface {
	if s.unitsService != nil {
		return s.unitsService
	}

	s.unitsService = &UnitsService{
		service: s,
	}

	return s.unitsService
}

// Users returns usersService
func (s *Service) Users() UsersServiceInterface {
	if s.usersService != nil {
		return s.usersService
	}

	s.usersService = &UsersService{
		service: s,
	}

	return s.usersService
}

// Variations returns VariationsService
func (s *Service) Variations() VariationsServiceInterface {
	if s.variationsService != nil {
		return s.variationsService
	}

	s.variationsService = &VariationsService{
		service: s,
	}

	return s.variationsService
}
