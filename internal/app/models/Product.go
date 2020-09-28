package models

import "encoding/json"

// Product ...
type Product struct {
	ID            int            `json:"id"`
	Name          string         `json:"name,omitempty"`
	Category      []*Category    `json:"categories,omitemty"`
	Options       []*Option      `json:"options,omitempty"`
	OptionsValues []*OptionValue `json:"options_values,omitempty"`
	Variations    []*Variation   `json:"variations,omitempty"`
}

// Validate ...
func (s *Product) Validate() error {
	return nil
}

// NewProductFromByte creates struct from byte slice
func NewProductFromByte(b []byte) (*Product, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &Product{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewProductSliceFromByte creates struct from byte slice
func NewProductSliceFromByte(b []byte) ([]*Product, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*Product = make([]*Product, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
