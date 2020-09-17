package models

import validation "github.com/go-ozzo/ozzo-validation"

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
