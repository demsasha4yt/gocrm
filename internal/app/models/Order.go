package models

// Order ...
type Order struct {
	ID           int       `json:"id"`
	Customer     *Customer `json:"customer,omitempty"`
	Address      string    `json:"address,omitempty"`
	Phone        string    `json:"phone,omitemty"`
	Email        string    `json:"email,omitempty"`
	ShippingDate string    `json:"shipping_date,omitempty"`
	Unit         *Unit     `json:"unit,omitempty"`
}

// Validate ...
func (s *Order) Validate() error {
	return nil
}
