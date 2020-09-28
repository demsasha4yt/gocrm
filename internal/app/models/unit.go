package models

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Unit ...
type Unit struct {
	ID            int             `json:"id"`
	Name          string          `json:"name,omitempty"`
	Address       string          `json:"address,omitempty"`
	Users         []*User         `json:"users,omitempty"`
	Manufacturers []*Manufacturer `json:"manufacturers,omitempty"`
	Categories    []*Manufacturer `json:"categories,omitempty"`
}

// Validate validates unit
func (u *Unit) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.ID, validation.Min(1)),
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Address, validation.Required),
	)
}

// NewUnitFromByte creates struct from byte slice
func NewUnitFromByte(b []byte) (*Unit, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &Unit{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewUnitSliceFromByte creates struct from byte slice
func NewUnitSliceFromByte(b []byte) ([]*Unit, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*Unit = make([]*Unit, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
