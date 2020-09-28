package models

import (
	"encoding/json"

	"github.com/guregu/null"
)

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

// NewCategoryFromByte creates struct from byte slice
func NewCategoryFromByte(b []byte) (*Category, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &Category{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewCategorySliceFromByte creates struct from byte slice
func NewCategorySliceFromByte(b []byte) ([]*Category, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*Category = make([]*Category, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
