package models

import "encoding/json"

// Option ...
type Option struct {
	ID           int         `json:"id"`
	Name         string      `json:"name,omitempty"`
	Description  string      `json:"description,omitempty"`
	OptionTypeID int         `json:"type_id,omitempty"`
	Type         *OptionType `json:"type,omitempty"`
}

// Validate ...
func (s *Option) Validate() error {
	return nil
}

// NewOptionFromByte creates struct from byte slice
func NewOptionFromByte(b []byte) (*Option, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	u := &Option{}
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewOptionSliceFromByte creates struct from byte slice
func NewOptionSliceFromByte(b []byte) ([]*Option, error) {
	if b == nil {
		return nil, ErrByteSliceNil
	}
	var u []*Option = make([]*Option, 0)
	if err := json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
