package models

// OptionType ...
type OptionType struct {
	ID     int    `json:"id"`
	Name   string `json:"name,omitempty"`
	IsSoft bool   `json:"is_soft"`
}

// Validate ...
func (s *OptionType) Validate() error {
	return nil
}
