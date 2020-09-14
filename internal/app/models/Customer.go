package models

// Customer ...
type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Address   string `json:"address,omitempty"`
	CreatedBy *User  `json:"created_by,omitempty"`
}

// Validate ...
func (s *Customer) Validate() error {
	return nil
}
