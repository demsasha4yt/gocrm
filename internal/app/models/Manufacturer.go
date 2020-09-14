package models

// Manufacturer ...
type Manufacturer struct {
	ID          int     `json:"id"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Units       []*Unit `json:"units,omitempty"`
}

// Validate ...
func (s *Manufacturer) Validate() error {
	return nil
}
