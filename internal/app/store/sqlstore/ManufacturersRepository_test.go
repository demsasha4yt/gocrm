package sqlstore_test

import (
	"fmt"
	"testing"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestManufacturersRepository_Create(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("manufacturers")

	s := sqlstore.New(db)
	m := models.TestManufacturer(t)

	assert.NoError(t, s.Manufacturers().Create(m))
	assert.NotEqual(t, m.ID, 0)

	m2 := &models.Manufacturer{
		Name: "HAHA",
	}
	assert.NoError(t, s.Manufacturers().Create(m2))
	assert.NotEqual(t, m2.ID, 0)
}

func TestManufacturers_FindAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	m := models.TestManufacturer(t)

	assert.NoError(t, s.Manufacturers().Create(m))
	assert.NotEqual(t, m.ID, 0)

	r, err := s.Manufacturers().FindAll()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestManufacturersRepository_Find(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("manufacturers")

	s := sqlstore.New(db)
	m := models.TestManufacturer(t)

	assert.NoError(t, s.Manufacturers().Create(m))
	assert.NotEqual(t, m.ID, 0)

	m2 := &models.Manufacturer{
		Name: "HAHA",
	}
	assert.NoError(t, s.Manufacturers().Create(m2))
	assert.NotEqual(t, m2.ID, 0)

	f, err := s.Manufacturers().Find(m.ID)
	assert.NoError(t, err)
	assert.NotNil(t, f)
	f2, err := s.Manufacturers().Find(m2.ID)
	assert.NoError(t, err)
	assert.NotNil(t, f2)

	f3, err := s.Manufacturers().Find(1000500)
	assert.Error(t, err)
	assert.Nil(t, f3)
}

func TestManufacturersRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("manufacturers")

	s := sqlstore.New(db)
	m := models.TestManufacturer(t)

	assert.NoError(t, s.Manufacturers().Create(m))
	assert.NotEqual(t, m.ID, 0)

	assert.NoError(t, s.Manufacturers().Delete(m.ID))

	f3, err := s.Manufacturers().Find(m.ID)
	assert.Error(t, err)
	assert.Nil(t, f3)
}

func TestManufacturersRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown()

	s := sqlstore.New(db)

	m := models.TestManufacturer(t)

	assert.NoError(t, s.Manufacturers().Create(m))
	assert.NotEqual(t, m.ID, 0)
	fmt.Printf("%d\n", m.ID)
	m2 := &models.Manufacturer{
		Name:        "Hello guys",
		Description: "Hello guys",
	}

	assert.NoError(t, s.Manufacturers().Update(m.ID, m2))
	fmt.Printf("%d\n", m.ID)
	m3, err := s.Manufacturers().Find(m.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Hello guys", m3.Name)
	assert.Equal(t, "Hello guys", m3.Description)
}
