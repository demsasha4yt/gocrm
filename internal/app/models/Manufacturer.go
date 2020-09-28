package models

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Manufacturer ...
type Manufacturer struct {
	ID          int     `json:"id"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Units       []*Unit `json:"units,omitempty"`
}

// Validate ...
func (s *Manufacturer) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required, validation.Length(1, 64)),
		validation.Field(&s.Description, validation.Length(1, 1024)),
	)
}

// NewManufacturerFromByte creates struct from byte slice
func NewManufacturerFromByte(b []byte) (*Manufacturer, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &Manufacturer{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewManufacturerSliceFromByte creates struct from byte slice
func NewManufacturerSliceFromByte(b []byte) ([]*Manufacturer, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*Manufacturer = make([]*Manufacturer, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
