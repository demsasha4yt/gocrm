package sqlstore_test

import (
	"context"
	"testing"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
	"github.com/demsasha4yt/gocrm.git/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func initSoftCategories(t *testing.T) (store.Store, func(...string)) {
	t.Helper()
	db, teardown := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	return s, teardown
}

func TestSoftCategoriesRepository_Create(t *testing.T) {
	s, teardown := initSoftCategories(t)
	defer teardown("softs_categories")

	u := &models.SoftCategory{
		Name:  "Lol",
		Value: 1,
	}

	assert.NoError(t, s.SoftCategories().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)
}

func TestSoftCategoriesRepository_Find(t *testing.T) {
	s, teardown := initSoftCategories(t)
	defer teardown("softs_categories")

	u := &models.SoftCategory{
		Name:  "Lol",
		Value: 1,
	}

	assert.NoError(t, s.SoftCategories().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	found, err := s.SoftCategories().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, u.Name, found.Name)
	assert.Equal(t, u.Value, found.Value)
}

func TestSoftCategoriesRepository_FindAll(t *testing.T) {
	s, teardown := initSoftCategories(t)
	defer teardown("softs_categories")

	assert.NoError(t, s.SoftCategories().Create(context.Background(), &models.SoftCategory{
		Name:  "Lol",
		Value: 1,
	}))
	assert.NoError(t, s.SoftCategories().Create(context.Background(), &models.SoftCategory{
		Name:  "Lol2",
		Value: 2,
	}))

	found, err := s.SoftCategories().FindAll(context.Background(), 0, 25)
	assert.NoError(t, err)
	assert.Len(t, found, 2)
}

func TestSoftCategoriesRepository_Delete(t *testing.T) {
	s, teardown := initSoftCategories(t)
	defer teardown("softs_categories")

	u := &models.SoftCategory{
		Name:  "Lol",
		Value: 1,
	}

	assert.NoError(t, s.SoftCategories().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	assert.NoError(t, s.SoftCategories().Delete(context.Background(), u.ID))

	found, err := s.SoftCategories().Find(context.Background(), u.ID)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestSoftCategoriesRepository_Update(t *testing.T) {
	s, teardown := initSoftCategories(t)
	defer teardown("softs_categories")

	u := &models.SoftCategory{
		Name:  "Lol",
		Value: 1,
	}

	assert.NoError(t, s.SoftCategories().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	update := &models.SoftCategory{
		Name:  "Bar",
		Value: 100,
	}

	assert.NoError(t, s.SoftCategories().Update(context.Background(), u.ID, update))

	found, err := s.SoftCategories().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, "Bar", found.Name)
	assert.Equal(t, 100, found.Value)
}
