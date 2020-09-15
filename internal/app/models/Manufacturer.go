package models

import (
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
		validation.Field(&s.Units),
	)
}
