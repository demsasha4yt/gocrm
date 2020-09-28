package models

import "encoding/json"

// Variation ...
type Variation struct {
	ID            int            `json:"id"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitemty"`
	Price         int            `json:"price,omitempty"`
	ProductID     int            `json:"product_id,omitempty"`
	Options       []*Option      `json:"options,omitempty"`
	OptionsValues []*OptionValue `json:"options_values,omitempty"`
	Product       *Product       `json:"product,omitempty"`
}

// Validate ...
func (s *Variation) Validate() error {
	return nil
}

// NewVariationFromByte creates struct from byte slice
func NewVariationFromByte(b []byte) (*Variation, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &Variation{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewVariationSliceFromByte creates struct from byte slice
func NewVariationSliceFromByte(b []byte) ([]*Variation, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*Variation = make([]*Variation, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
