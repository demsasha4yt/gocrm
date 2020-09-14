package models

// Category ...
type Category struct {
	ID          int        `json:"id"`
	Name        string     `json:"name,omitempty"`
	Descriprion string     `json:"description,omitempty"`
	IsSoft      bool       `json:"is_soft"`
	Parent      *Category  `json:"parent,omitempty"`
	Products    []*Product `json:"products,omitempty"`
}

// Validate ...
func (s *Category) Validate() error {
	return nil
}
