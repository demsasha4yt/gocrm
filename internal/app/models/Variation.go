package models

// Variation ...
type Variation struct {
	ID            int            `json:"id"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitemty"`
	Options       []*Option      `json:"options,omitempty"`
	OptionsValues []*OptionValue `json:"options_values,omitempty"`
	Product       *Product       `json:"product,omitempty"`
}

// Validate ...
func (s *Variation) Validate() error {
	return nil
}
