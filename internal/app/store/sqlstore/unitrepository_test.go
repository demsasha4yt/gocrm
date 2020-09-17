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
	assert.NoError(t, s.Unit().Create(u))
	assert.NotNil(t, u.ID)
}

func TestUnitRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("units")

	s := sqlstore.New(db)

	u := models.TestUnit(t)
	assert.NoError(t, s.Unit().Create(u))
	assert.NotNil(t, u.ID)
	unit, err := s.Unit().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, unit)
	assert.Equal(t, u.ID, unit.ID)
}

func TestUnitRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("units")

	s := sqlstore.New(db)
	u := models.TestUnit(t)
	assert.NoError(t, s.Unit().Create(u))
	assert.NotNil(t, u.ID)

	assert.NoError(t, s.Unit().Delete(u.ID))
	unit, err := s.Unit().Find(100500)
	assert.Error(t, err)
	assert.Nil(t, unit)
}

func TestUnitRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("units")

	s := sqlstore.New(db)
	u := models.TestUnit(t)
	assert.NoError(t, s.Unit().Create(u))
	assert.NotNil(t, u.ID)

	newU := &models.Unit{
		Name:    "New Name",
		Address: "New Address",
	}

	assert.NoError(t, s.Unit().Update(u.ID, newU))
	assert.Equal(t, "New Name", newU.Name)
	assert.Equal(t, "New Address", newU.Address)

	unit, err := s.Unit().Find(u.ID)
	assert.NoError(t, err)
	assert.Equal(t, "New Name", unit.Name)
	assert.Equal(t, "New Address", unit.Address)

	assert.Error(t, s.Unit().Update(100500, newU))
}

func TestUnitRepository_FindUnitsByUserID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("units", "users")

	s := sqlstore.New(db)
	u := models.TestUnit(t)
	u.Name = "1"
	u2 := models.TestUnit(t)
	u2.Name = "2"
	assert.NoError(t, s.Unit().Create(u))
	assert.NoError(t, s.Unit().Create(u2))

	user := models.TestUser(t)
	assert.NoError(t, s.User().Create(user))
	if _, err := db.Query(context.Background(), `INSERT INTO users_units(user_id, unit_id)
	VALUES ($1, $2), ($1, $3)`, user.ID, u.ID, u2.ID); err != nil {
		t.Fatal(err)
	}
	units, err := s.Unit().FindUnitsByUserID(user.ID)
	assert.NoError(t, err)
	assert.Len(t, units, 2)

	user2 := models.TestUser(t)
	user2.Login = "2"
	user2.Email = "2@yandex.ru"
	assert.NoError(t, s.User().Create(user2))
	units, err = s.Unit().FindUnitsByUserID(user2.ID)
	assert.NoError(t, err)
	assert.Len(t, units, 0)
}
