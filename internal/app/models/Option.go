package models

// Option ...
type Option struct {
	ID          int         `json:"id"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Type        *OptionType `json:"type,omitempty"`
}

// Validate ...
func (s *Option) Validate() error {
	return nil
}
