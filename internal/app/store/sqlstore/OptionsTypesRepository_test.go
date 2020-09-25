package sqlstore_test

import (
	"context"
	"testing"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store"
	"github.com/demsasha4yt/gocrm.git/internal/app/store/sqlstore"
	"github.com/guregu/null"
	"github.com/stretchr/testify/assert"
)

func initStore(t *testing.T) (store.Store, func(...string)) {
	t.Helper()
	db, teardown := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	return s, teardown
}

func TestOptionsTypesRepository_Create(t *testing.T) {
	s, teardown := initStore(t)
	defer teardown("options_types")

	u := &models.OptionType{
		Name:   "Hello",
		IsSoft: null.NewBool(false, true),
	}
	u2 := &models.OptionType{
		Name:   "Hello2",
		IsSoft: null.NewBool(true, true),
	}
	assert.NoError(t, s.OptionsTypes().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)
	assert.NoError(t, s.OptionsTypes().Create(context.Background(), u2))
	assert.NotEqual(t, u2.ID, 0)
}

func TestOptionsTypesRepository_Find(t *testing.T) {
	s, teardown := initStore(t)
	defer teardown("options_types")

	u := &models.OptionType{
		Name:   "Hello",
		IsSoft: null.NewBool(false, true),
	}
	assert.NoError(t, s.OptionsTypes().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)
	found, err := s.OptionsTypes().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	found2, err2 := s.OptionsTypes().Find(context.Background(), 100500)
	assert.Error(t, err2)
	assert.Nil(t, found2)
}

func TestOptionsTypesRepository_FindAll(t *testing.T) {
	s, teardown := initStore(t)
	defer teardown("options_types")

	a, err := s.OptionsTypes().FindAll(context.Background(), 0, 25)
	assert.NoError(t, err)
	assert.Len(t, a, 0)

	assert.NoError(t, s.OptionsTypes().Create(context.Background(), &models.OptionType{
		Name:   "Hello",
		IsSoft: null.NewBool(false, true),
	}))
	a, err = s.OptionsTypes().FindAll(context.Background(), 0, 25)
	assert.NoError(t, err)
	assert.Len(t, a, 1)

	a, err = s.OptionsTypes().FindAll(context.Background(), 0, 0)
	assert.NoError(t, err)
	assert.Len(t, a, 0)

	a, err = s.OptionsTypes().FindAll(context.Background(), 1, 25)
	assert.NoError(t, err)
	assert.Len(t, a, 0)
}

func TestOptionsTypesRepository_Delete(t *testing.T) {
	s, teardown := initStore(t)
	defer teardown("options_types")

	u := &models.OptionType{
		Name:   "Hello",
		IsSoft: null.NewBool(false, true),
	}
	assert.NoError(t, s.OptionsTypes().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	found, err := s.OptionsTypes().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	assert.NoError(t, s.OptionsTypes().Delete(context.Background(), u.ID))

	found, err = s.OptionsTypes().Find(context.Background(), u.ID)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestOptionsTypesRepository_Update(t *testing.T) {
	s, teardown := initStore(t)
	defer teardown("options_types")

	u := &models.OptionType{
		Name:   "Hello",
		IsSoft: null.NewBool(true, true),
	}
	assert.NoError(t, s.OptionsTypes().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	update := &models.OptionType{
		Name: "Hello2",
	}
	assert.NoError(t, s.OptionsTypes().Update(context.Background(), u.ID, update))
	found, err := s.OptionsTypes().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, "Hello2", found.Name)
	assert.True(t, found.IsSoft.Equal(null.BoolFrom(true)))

	update = &models.OptionType{
		IsSoft: null.BoolFrom(false),
	}
	assert.NoError(t, s.OptionsTypes().Update(context.Background(), u.ID, update))
	found, err = s.OptionsTypes().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.True(t, found.IsSoft.Equal(null.BoolFrom(false)))
}
