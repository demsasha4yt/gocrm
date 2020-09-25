package models

import "github.com/guregu/null"

// OptionType ...
type OptionType struct {
	ID     int       `json:"id"`
	Name   string    `json:"name,omitempty"`
	IsSoft null.Bool `json:"is_soft"`
}

// Validate ...
func (s *OptionType) Validate() error {
	return nil
}
