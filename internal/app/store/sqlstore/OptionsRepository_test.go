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

func initOptions(t *testing.T) (store.Store, func(...string), int) {
	t.Helper()
	db, teardown := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	u := &models.OptionType{
		Name:   "Hello",
		IsSoft: null.NewBool(false, true),
	}
	s.OptionsTypes().Create(context.Background(), u)
	return s, teardown, u.ID
}

func TestOptionsRepository_Create(t *testing.T) {
	s, teardown, tid := initOptions(t)
	defer teardown("options_types", "options")

	u := &models.Option{
		Name:         "Кикек",
		Description:  "Кек",
		OptionTypeID: tid,
	}
	assert.NoError(t, s.Options().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)
}

func TestOptionsRepository_Find(t *testing.T) {
	s, teardown, tid := initOptions(t)
	defer teardown("options_types", "options")

	u := &models.Option{
		Name:         "Кикек",
		Description:  "Кек",
		OptionTypeID: tid,
	}
	assert.NoError(t, s.Options().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	found, err := s.Options().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)

	found, err = s.Options().Find(context.Background(), 105000)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestOptionsRepository_FindAll(t *testing.T) {
	s, teardown, tid := initOptions(t)
	defer teardown("options_types", "options")

	found, err := s.Options().FindAll(context.Background(), 0, 25)
	assert.NoError(t, err)
	assert.Len(t, found, 0)

	assert.NoError(t, s.Options().Create(context.Background(), &models.Option{
		Name:         "Кикек",
		Description:  "Кек",
		OptionTypeID: tid,
	}))
	found, err = s.Options().FindAll(context.Background(), 0, 25)
	assert.NoError(t, err)
	assert.Len(t, found, 1)
	found, err = s.Options().FindAll(context.Background(), 0, 0)
	assert.NoError(t, err)
	assert.Len(t, found, 0)
	found, err = s.Options().FindAll(context.Background(), 1, 25)
	assert.NoError(t, err)
	assert.Len(t, found, 0)
}

func TestOptionsRepository_Delete(t *testing.T) {
	s, teardown, tid := initOptions(t)
	defer teardown("options_types", "options")

	u := &models.Option{
		Name:         "Кикек",
		Description:  "Кек",
		OptionTypeID: tid,
	}
	assert.NoError(t, s.Options().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	assert.NoError(t, s.Options().Delete(context.Background(), u.ID))
	found, err := s.Options().Find(context.Background(), u.ID)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestOptionsRepository_Update(t *testing.T) {
	s, teardown, tid := initOptions(t)
	defer teardown("options_types", "options")

	u := &models.Option{
		Name:         "Кикек",
		Description:  "Кек",
		OptionTypeID: tid,
	}
	assert.NoError(t, s.Options().Create(context.Background(), u))
	assert.NotEqual(t, u.ID, 0)

	update := &models.Option{
		Name:        "Bar",
		Description: "Desc",
	}

	assert.NoError(t, s.Options().Update(context.Background(), u.ID, update))

	found, err := s.Options().Find(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, "Bar", found.Name)
	assert.Equal(t, "Desc", found.Description)
	assert.Equal(t, u.OptionTypeID, found.OptionTypeID)
}
