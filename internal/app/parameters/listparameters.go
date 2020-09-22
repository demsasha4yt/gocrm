package listparameters

// ListParameters ...
type ListParameters interface{}

//PaginationParameters ...
type PaginationParameters struct {
	Page     int
	PageSize int
}

// Parameters ...
type Parameters struct {
	ListParameters
	*PaginationParameters
}
