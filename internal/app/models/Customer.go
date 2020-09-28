package models

import "encoding/json"

// Customer ...
type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Address   string `json:"address,omitempty"`
	CreatedBy *User  `json:"created_by,omitempty"`
}

// Validate ...
func (s *Customer) Validate() error {
	return nil
}

// NewCustomerFromByte creates struct from byte slice
func NewCustomerFromByte(b []byte) (*Customer, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &Customer{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewCustomerSliceFromByte creates struct from byte slice
func NewCustomerSliceFromByte(b []byte) ([]*Customer, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*Customer = make([]*Customer, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
