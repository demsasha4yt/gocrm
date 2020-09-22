package sqlstore_test

import (
	"context"
	"testing"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store/sqlstore"
	"github.com/guregu/null"
	"github.com/stretchr/testify/assert"
)

func TestCategoriesRepository_Create(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	m := models.TestCategory(t)

	assert.NoError(t, s.Categories().Create(context.Background(), m))
	assert.NotEqual(t, m.ID, 0)

	m2 := &models.Category{
		Name:     "HAHA",
		ParentID: null.NewInt(int64(m.ID), true),
	}
	assert.NoError(t, s.Categories().Create(context.Background(), m2))
	assert.NotEqual(t, m2.ID, 0)
}

func TestCategories_FindAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	m := models.TestCategory(t)

	assert.NoError(t, s.Categories().Create(context.Background(), m))
	assert.NotEqual(t, m.ID, 0)

	r, err := s.Categories().FindAll(context.Background(), 0, 1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestCategoriesRepository_Find(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	m := models.TestCategory(t)

	assert.NoError(t, s.Categories().Create(context.Background(), m))
	assert.NotEqual(t, m.ID, 0)

	m2 := &models.Category{
		Name: "HAHA",
	}
	assert.NoError(t, s.Categories().Create(context.Background(), m2))
	assert.NotEqual(t, m2.ID, 0)

	f, err := s.Categories().Find(context.Background(), m.ID)
	assert.NoError(t, err)
	assert.NotNil(t, f)
	f2, err := s.Categories().Find(context.Background(), m2.ID)
	assert.NoError(t, err)
	assert.NotNil(t, f2)

	f3, err := s.Categories().Find(context.Background(), 1000500)
	assert.Error(t, err)
	assert.Nil(t, f3)
}

func TestCategoriesRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	m := models.TestCategory(t)

	assert.NoError(t, s.Categories().Create(context.Background(), m))
	assert.NotEqual(t, m.ID, 0)

	assert.NoError(t, s.Categories().Delete(context.Background(), m.ID))

	f3, err := s.Categories().Find(context.Background(), m.ID)
	assert.Error(t, err)
	assert.Nil(t, f3)
}

func TestCategoriesRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)

	m := models.TestCategory(t)

	assert.NoError(t, s.Categories().Create(context.Background(), m))
	assert.NotEqual(t, m.ID, 0)

	m2 := &models.Category{
		Name:        "Hello guys",
		Description: "Hello guys",
	}

	assert.NoError(t, s.Categories().Update(context.Background(), m.ID, m2))

	m3, err := s.Categories().Find(context.Background(), m.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Hello guys", m3.Name)
	assert.Equal(t, "Hello guys", m3.Description)
}
