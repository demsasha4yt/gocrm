package models

import (
	"encoding/json"

	"github.com/guregu/null"
)

// OptionType ...
type OptionType struct {
	ID     int       `json:"id"`
	Name   string    `json:"name,omitempty"`
	IsSoft null.Bool `json:"is_soft"`
}

// Validate ...
func (s *OptionType) Validate() error {
	return nil
}

// NewOptionTypeFromByte creates struct from byte slice
func NewOptionTypeFromByte(b []byte) (*OptionType, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &OptionType{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewOptionTypeSliceFromByte creates struct from byte slice
func NewOptionTypeSliceFromByte(b []byte) ([]*OptionType, error) {
	var u []*OptionType = make([]*OptionType, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
