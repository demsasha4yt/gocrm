package models

// OptionValue ...
type OptionValue struct {
	ID     int           `json:"id"`
	Value  string        `json:"value,omitempty"`
	Image  string        `json:"image,omitempty"`
	Option *Option       `json:"option,omitempty"`
	Type   *OptionType   `json:"type,omitempty"`
	Softs  []*OptionSoft `json:"softs,omitempty"`
}

// Validate ...
func (s *OptionValue) Validate() error {
	return nil
}
