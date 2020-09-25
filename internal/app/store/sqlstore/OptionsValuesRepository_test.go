package sqlstore_test

import (
	"testing"

	"github.com/demsasha4yt/gocrm.git/internal/app/store"
	"github.com/demsasha4yt/gocrm.git/internal/app/store/sqlstore"
)

func initOptionsValues(t *testing.T) (store.Store, func(...string)) {
	t.Helper()
	db, teardown := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	return s, teardown
}

func TestOptionsValuesRepository_Create(t *testing.T) {

}

func TestOptionsValuesRepository_Find(t *testing.T) {

}

func TestOptionsValuesRepository_FindAll(t *testing.T) {

}

func TestOptionsValuesRepository_Delete(t *testing.T) {

}

func TestOptionsValuesRepository_Update(t *testing.T) {

}
