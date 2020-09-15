package models_test

import (
	"testing"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestManufacturer_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *models.Manufacturer
		isValid bool
	}{
		{
			name: "Valid",
			s: func() *models.Manufacturer {
				return models.TestManufacturer(t)
			},
			isValid: true,
		},
		{
			name: "Long name",
			s: func() *models.Manufacturer {
				s := models.TestManufacturer(t)
				s.Name = `sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss
				sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss
				ssssssssssssssssssssssssssssssssssssssssssssssss`
				return s
			},
			isValid: false,
		},
		{
			name: "Has no units",
			s: func() *models.Manufacturer {
				s := models.TestManufacturer(t)
				s.Units = nil
				return s
			},
			isValid: true,
		},
		{
			name: "Has no name",
			s: func() *models.Manufacturer {
				s := models.TestManufacturer(t)
				s.Name = ""
				return s
			},
			isValid: false,
		},
		{
			name: "Has no description",
			s: func() *models.Manufacturer {
				s := models.TestManufacturer(t)
				s.Description = ""
				return s
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().Validate())
			} else {
				assert.Error(t, tc.s().Validate())
			}
		})
	}
}
