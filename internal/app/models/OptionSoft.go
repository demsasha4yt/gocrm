package models

// OptionSoft ...
type OptionSoft struct {
	ID           int           `json:"id"`
	Name         string        `json:"name,omitempty"`
	Image        string        `json:"image,omitempty"`
	IsSoft       bool          `json:"is_soft"`
	OptionValue  *OptionValue  `json:"option_value,omitempty"`
	Manufacturer *Manufacturer `json:"manufacturer,omitempty"`
}

// Validate ...
func (s *OptionSoft) Validate() error {
	return nil
}
