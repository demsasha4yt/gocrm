package models

import "github.com/guregu/null"

// Category ...
type Category struct {
	ID            int         `json:"id"`
	Name          string      `json:"name,omitempty"`
	Description   string      `json:"description,omitempty"`
	ParentID      null.Int    `json:"parentid,omitempty"`
	Subcategories []*Category `json:"subcategories,omitempty"`
	Products      []*Product  `json:"products,omitempty"`
	Units         []*Unit     `json:"units,omitempty"`
}

// Validate ...
func (s *Category) Validate() error {
	return nil
}
