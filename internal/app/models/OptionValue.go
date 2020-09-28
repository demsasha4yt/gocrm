package models

import "encoding/json"

// OptionValue ...
type OptionValue struct {
	ID           int           `json:"id"`
	Value        string        `json:"value,omitempty"`
	Image        string        `json:"image,omitempty"`
	OptionID     int           `json:"option_id,omitempty"`
	OptionTypeID int           `json:"option_type_id,omitempty"`
	Option       *Option       `json:"option,omitempty"`
	OptionType   *OptionType   `json:"type,omitempty"`
	Softs        []*OptionSoft `json:"softs,omitempty"`
}

// Validate ...
func (s *OptionValue) Validate() error {
	return nil
}

// NewOptionValueFromByte creates struct from byte slice
func NewOptionValueFromByte(b []byte) (*OptionValue, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &OptionValue{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewOptionValueSliceFromByte creates struct from byte slice
func NewOptionValueSliceFromByte(b []byte) ([]*OptionValue, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*OptionValue = make([]*OptionValue, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
