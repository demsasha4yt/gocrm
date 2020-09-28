package models

import "encoding/json"

// Order ...
type Order struct {
	ID           int    `json:"id"`
	Order        *Order `json:"Order,omitempty"`
	Address      string `json:"address,omitempty"`
	Phone        string `json:"phone,omitemty"`
	Email        string `json:"email,omitempty"`
	ShippingDate string `json:"shipping_date,omitempty"`
	Unit         *Unit  `json:"unit,omitempty"`
}

// Validate ...
func (s *Order) Validate() error {
	return nil
}

// NewOrderFromByte creates struct from byte slice
func NewOrderFromByte(b []byte) (*Order, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &Order{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewOrderSliceFromByte creates struct from byte slice
func NewOrderSliceFromByte(b []byte) ([]*Order, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*Order = make([]*Order, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
