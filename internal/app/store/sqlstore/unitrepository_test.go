package sqlstore_test

import (
	"context"
	"testing"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUnitRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("units")

	s := sqlstore.New(db)
	u := models.TestUnit(t)
	assert.NoError(t, s.Unit().Create(context.Background(), u))
	assert.NotNil(t, u.ID)
}

func TestUnitRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("units")

	s := sqlstore.New(db)

	u := models.TestUnit(t)
	assert.NoError(t, s.Unit().Create(context.Background(), u))
	assert.NotNil(t, u.ID)
	unit, err := s.Unit().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, unit)
	assert.Equal(t, u.ID, unit.ID)
}

func TestUnitRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("units")

	s := sqlstore.New(db)
	u := models.TestUnit(t)
	assert.NoError(t, s.Unit().Create(context.Background(), u))
	assert.NotNil(t, u.ID)

	assert.NoError(t, s.Unit().Delete(context.Background(), u.ID))
	unit, err := s.Unit().Find(context.Background(), 100500)
	assert.Error(t, err)
	assert.Nil(t, unit)
}

func TestUnitRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("units")

	s := sqlstore.New(db)
	u := models.TestUnit(t)
	assert.NoError(t, s.Unit().Create(context.Background(), u))
	assert.NotNil(t, u.ID)

	newU := &models.Unit{
		Name:    "New Name",
		Address: "New Address",
	}

	assert.NoError(t, s.Unit().Update(context.Background(), u.ID, newU))
	assert.Equal(t, "New Name", newU.Name)
	assert.Equal(t, "New Address", newU.Address)

	unit, err := s.Unit().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.Equal(t, "New Name", unit.Name)
	assert.Equal(t, "New Address", unit.Address)

	assert.Error(t, s.Unit().Update(context.Background(), 100500, newU))
}
