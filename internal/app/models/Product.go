package models

// Product ...
type Product struct {
	ID            int            `json:"id"`
	Name          string         `json:"name,omitempty"`
	Category      []*Category    `json:"categories,omitemty"`
	Options       []*Option      `json:"options,omitempty"`
	OptionsValues []*OptionValue `json:"options_values,omitempty"`
	Variations    []*Variation   `json:"variations,omitempty"`
}

// Validate ...
func (s *Product) Validate() error {
	return nil
}
