package sqlstore_test

import (
	"context"
	"testing"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func initOptionsValues(t *testing.T) (store.Store, func(...string), int, int) {
	t.Helper()
	s, teardown, optionTypeID := initOptions(t)
	u := &models.Option{
		Name:         "Option",
		Description:  "Option desc",
		OptionTypeID: optionTypeID,
	}
	s.Options().Create(context.Background(), u)
	return s, teardown, optionTypeID, u.ID
}

func TestOptionsValuesRepository_Create(t *testing.T) {
	s, teardown, optionTypeID, optionID := initOptionsValues(t)
	defer teardown("options_values", "options_types", "options")
	u := &models.OptionValue{
		Value:        "100",
		Image:        "Image",
		OptionID:     optionID,
		OptionTypeID: optionTypeID,
	}
	assert.NoError(t, s.OptionsValues().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)
}

func TestOptionsValuesRepository_Find(t *testing.T) {
	s, teardown, optionTypeID, optionID := initOptionsValues(t)
	defer teardown("options_values", "options_types", "options")
	u := &models.OptionValue{
		Value:        "100",
		Image:        "Image",
		OptionID:     optionID,
		OptionTypeID: optionTypeID,
	}
	assert.NoError(t, s.OptionsValues().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)
	found, err := s.OptionsValues().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	found, err = s.OptionsValues().Find(context.Background(), 1005000)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestOptionsValuesRepository_FindAll(t *testing.T) {
	s, teardown, optionTypeID, optionID := initOptionsValues(t)
	defer teardown("options_values", "options_types", "options")
	found, err := s.OptionsValues().FindAll(context.Background(), 0, 100)
	assert.NoError(t, err)
	assert.Len(t, found, 0)
	assert.NoError(t, s.OptionsValues().Create(context.Background(), &models.OptionValue{
		Value:        "100",
		Image:        "Image",
		OptionID:     optionID,
		OptionTypeID: optionTypeID,
	}))
	found, err = s.OptionsValues().FindAll(context.Background(), 0, 100)
	assert.NoError(t, err)
	assert.Len(t, found, 1)
}

func TestOptionsValuesRepository_Delete(t *testing.T) {
	s, teardown, optionTypeID, optionID := initOptionsValues(t)
	defer teardown("options_values", "options_types", "options")
	u := &models.OptionValue{
		Value:        "100",
		Image:        "Image",
		OptionID:     optionID,
		OptionTypeID: optionTypeID,
	}
	assert.NoError(t, s.OptionsValues().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	assert.NoError(t, s.OptionsValues().Delete(context.Background(), u.ID))

	found, err := s.OptionsValues().Find(context.Background(), u.ID)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestOptionsValuesRepository_Update(t *testing.T) {
	s, teardown, optionTypeID, optionID := initOptionsValues(t)
	defer teardown("options_values", "options_types", "options")
	u := &models.OptionValue{
		Value:        "100",
		Image:        "Image",
		OptionID:     optionID,
		OptionTypeID: optionTypeID,
	}
	assert.NoError(t, s.OptionsValues().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	update := &models.OptionValue{
		Value:        "102",
		Image:        "Test",
		OptionID:     optionID,
		OptionTypeID: optionTypeID,
	}
	assert.NoError(t, s.OptionsValues().Update(context.Background(), u.ID, update))

	found, err := s.OptionsValues().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	assert.Equal(t, "Test", found.Image)
	assert.Equal(t, "102", found.Value)
	assert.Equal(t, optionID, found.OptionID)
	assert.Equal(t, optionTypeID, found.OptionTypeID)
}
