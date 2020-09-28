package models

import "encoding/json"

// OptionSoft ...
type OptionSoft struct {
	ID             int           `json:"id"`
	Name           string        `json:"name,omitempty"`
	Image          string        `json:"image,omitempty"`
	ManufacturerID int           `json:"manufacturer_id,omitempty"`
	OptionValueID  int           `json:"option_value_id,omitempty"`
	SoftCategoryID int           `json:"soft_category_id,omitempty"`
	SoftCategory   *SoftCategory `json:"soft_category,omitempty"`
	OptionValue    *OptionValue  `json:"option_value,omitempty"`
	Manufacturer   *Manufacturer `json:"manufacturer,omitempty"`
}

// Validate ...
func (s *OptionSoft) Validate() error {
	return nil
}

// NewOptionSoftFromByte creates struct from byte slice
func NewOptionSoftFromByte(b []byte) (*OptionSoft, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &OptionSoft{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewOptionSoftSliceFromByte creates struct from byte slice
func NewOptionSoftSliceFromByte(b []byte) ([]*OptionSoft, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*OptionSoft = make([]*OptionSoft, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
