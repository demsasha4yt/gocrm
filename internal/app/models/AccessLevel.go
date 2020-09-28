package models

import "encoding/json"

// AccessLevel ...
type AccessLevel struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	AccessCreate       bool   `json:"access_create"`
	AccessGet          bool   `json:"access_get"`
	AccessUpdate       bool   `json:"access_update"`
	AccessDelete       bool   `json:"access_delete"`
	CategoryCreate     bool   `json:"category_cretae"`
	CategoryGet        bool   `json:"category_get"`
	CategoryUpdate     bool   `json:"category_update"`
	CategoryDelete     bool   `json:"category_delete"`
	CustomerCreate     bool   `json:"customer_create"`
	CustomerGet        bool   `json:"customer_get"`
	CustomerUpdate     bool   `json:"customer_update"`
	CustomerDelete     bool   `json:"customer_delete"`
	ManufacturerCreate bool   `json:"manufacturer_create"`
	ManufacturerGet    bool   `json:"manufacturer_get"`
	ManufacturerUpdate bool   `json:"manufacturer_update"`
	ManufacturerDelete bool   `json:"manufacturer_delete"`
	OptionCreate       bool   `json:"option_create"`
	OptionGet          bool   `json:"option_get"`
	OptionUpdate       bool   `json:"option_update"`
	OptionDelete       bool   `json:"option_delete"`
	OrderCreate        bool   `json:"order_create"`
	OrderGet           bool   `json:"order_get"`
	OrderUpdate        bool   `json:"order_update"`
	OrderDelete        bool   `json:"order_delete"`
	ProductCreate      bool   `json:"product_create"`
	ProductGet         bool   `json:"product_get"`
	ProductUpdate      bool   `json:"prodct_update"`
	ProductDelete      bool   `json:"product_delete"`
	UnitCreate         bool   `json:"unit_create"`
	UnitGet            bool   `json:"unit_get"`
	UnitUpdate         bool   `json:"unit_update"`
	UnitDelete         bool   `json:"unit_delete"`
	UserCreate         bool   `json:"user_create"`
	UserGet            bool   `json:"user_get"`
	UserUpdate         bool   `json:"user_update"`
	UserDelete         bool   `json:"user_delete"`
}

// Validate ...
func (s *AccessLevel) Validate() error {
	return nil
}

// NewAccessLevelFromByte creates struct from byte slice
func NewAccessLevelFromByte(b []byte) (*AccessLevel, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &AccessLevel{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewAccessLevelSliceFromByte creates struct from byte slice
func NewAccessLevelSliceFromByte(b []byte) ([]*AccessLevel, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*AccessLevel = make([]*AccessLevel, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
