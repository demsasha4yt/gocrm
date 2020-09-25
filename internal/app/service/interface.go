package service

// Interface ...
type Interface interface {
	Categories() CategoriesServiceInterface
	Customers() CustomersServiceInterface
	Manufacturers() ManufacturersServiceInterface
	Options() OptionsServiceInterface
	OptionsSofts() OptionsSoftsServiceInterface
	OptionsTypes() OptionsTypesServiceInterface
	OptionsValues() OptionsValuesServiceInterface
	Orders() OrdersServiceInterface
	Units() UnitsServiceInterface
	Users() UsersServiceInterface
	Variations() VariationsServiceInterface
}
