package store

// Store interface
type Store interface {
	Categories() CategoriesRepository
	Customers() CustomersRepository
	Manufacturers() ManufacturersRepository
	Options() OptionsRepository
	OptionsSofts() OptionsSoftsRepository
	OptionsTypes() OptionsTypesRepository
	OptionsValues() OptionsValuesRepository
	Orders() OrdersRepository
	Unit() UnitRepository
	User() UserRepository
	Variations() VariationsRepository
}
