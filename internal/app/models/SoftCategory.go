package models

import "encoding/json"

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

// NewSoftCategoryFromByte creates struct from byte slice
func NewSoftCategoryFromByte(b []byte) (*SoftCategory, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &SoftCategory{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewSoftCategorySliceFromByte creates struct from byte slice
func NewSoftCategorySliceFromByte(b []byte) ([]*SoftCategory, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*SoftCategory = make([]*SoftCategory, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
