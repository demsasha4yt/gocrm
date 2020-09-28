package sqlstore_test

import (
	"context"
	"testing"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func initOptionsSofts(t *testing.T) (store.Store, func(...string), int, int, int, int, int) {
	s, teardown, optionTypeID, optionID := initOptionsValues(t)

	v := &models.OptionValue{
		Value:        "1000",
		OptionID:     optionID,
		OptionTypeID: optionTypeID,
	}

	s.OptionsValues().Create(context.Background(), v)

	u := &models.Manufacturer{
		Name:        "Manufacturer",
		Description: "Hello",
	}

	assert.NoError(t, s.Manufacturers().Create(context.Background(), u))

	g := &models.SoftCategory{
		Name:  "hello",
		Value: 7,
	}
	assert.NoError(t, s.SoftCategories().Create(context.Background(), g))
	return s, teardown, optionTypeID, optionID, v.ID, u.ID, g.ID
}

func TestOptionsSoftsRepository_Create(t *testing.T) {
	s, teardown, _, _, optionValueID, manufacturerID, softCategoryID := initOptionsSofts(t)
	defer teardown("options_values", "options_types", "options", "options_softs", "softs_categories", "manufacturers")
	u := &models.OptionSoft{
		Name:           "Soft",
		Image:          "Image",
		ManufacturerID: manufacturerID,
		OptionValueID:  optionValueID,
		SoftCategoryID: softCategoryID,
	}

	assert.NoError(t, s.OptionsSofts().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)
}

func TestOptionsSoftsRepository_Find(t *testing.T) {
	s, teardown, _, _, optionValueID, manufacturerID, softCategoryID := initOptionsSofts(t)
	defer teardown("options_values", "options_types", "options", "options_softs", "softs_categories", "manufacturers")
	u := &models.OptionSoft{
		Name:           "Soft",
		Image:          "Image",
		ManufacturerID: manufacturerID,
		OptionValueID:  optionValueID,
		SoftCategoryID: softCategoryID,
	}

	assert.NoError(t, s.OptionsSofts().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	found, err := s.OptionsSofts().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
}

func TestOptionsSoftsRepository_FindAll(t *testing.T) {
	s, teardown, _, _, optionValueID, manufacturerID, softCategoryID := initOptionsSofts(t)
	defer teardown("options_values", "options_types", "options", "options_softs", "softs_categories", "manufacturers")

	found, err := s.OptionsSofts().FindAll(context.Background(), 0, 100)
	assert.NoError(t, err)
	assert.Len(t, found, 0)

	assert.NoError(t, s.OptionsSofts().Create(context.Background(), &models.OptionSoft{
		Name:           "Soft",
		Image:          "Image",
		ManufacturerID: manufacturerID,
		OptionValueID:  optionValueID,
		SoftCategoryID: softCategoryID,
	}))

	found, err = s.OptionsSofts().FindAll(context.Background(), 0, 100)
	assert.NoError(t, err)
	assert.Len(t, found, 1)
}

func TestOptionsSoftsRepository_Delete(t *testing.T) {
	s, teardown, _, _, optionValueID, manufacturerID, softCategoryID := initOptionsSofts(t)
	defer teardown("options_values", "options_types", "options", "options_softs", "softs_categories", "manufacturers")
	u := &models.OptionSoft{
		Name:           "Soft",
		Image:          "Image",
		ManufacturerID: manufacturerID,
		OptionValueID:  optionValueID,
		SoftCategoryID: softCategoryID,
	}

	assert.NoError(t, s.OptionsSofts().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	assert.NoError(t, s.OptionsSofts().Delete(context.Background(), u.ID))

	found, err := s.OptionsSofts().Find(context.Background(), u.ID)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestOptionsSoftsRepository_Update(t *testing.T) {
	s, teardown, _, _, optionValueID, manufacturerID, softCategoryID := initOptionsSofts(t)
	defer teardown("options_values", "options_types", "options", "options_softs", "softs_categories", "manufacturers")
	u := &models.OptionSoft{
		Name:           "Soft",
		Image:          "Image",
		ManufacturerID: manufacturerID,
		OptionValueID:  optionValueID,
		SoftCategoryID: softCategoryID,
	}

	assert.NoError(t, s.OptionsSofts().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	update := &models.OptionSoft{
		Name:           "Soft2",
		Image:          "Image2",
		ManufacturerID: manufacturerID,
		OptionValueID:  optionValueID,
		SoftCategoryID: softCategoryID,
	}
	assert.NoError(t, s.OptionsSofts().Update(context.Background(), u.ID, update))

	found, err := s.OptionsSofts().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	assert.Equal(t, "Soft2", found.Name)
	assert.Equal(t, "Image2", found.Image)
	assert.Equal(t, u.ManufacturerID, found.ManufacturerID)
	assert.Equal(t, u.OptionValueID, found.OptionValueID)
	assert.Equal(t, u.SoftCategoryID, found.SoftCategoryID)
}
