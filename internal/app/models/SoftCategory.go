package models

// SoftCategory ...
type SoftCategory struct {
	ID    int    `json:"id"`
	Name  string `json:"name,omitempty"`
	Value int    `json:"value,omitempty"`
}

// Validate validates SoftCategory
func (u *SoftCategory) Validate() error {
	return nil
}
