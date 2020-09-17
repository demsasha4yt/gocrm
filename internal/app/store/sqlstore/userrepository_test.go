package sqlstore_test

import (
	"testing"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
	"github.com/demsasha4yt/gocrm.git/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := models.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := models.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := models.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.Error(t, err)

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := models.TestUser(t)
	_, err := s.User().FindByLogin(u1.Login)
	assert.Error(t, err)

	s.User().Create(u1)
	u2, err := s.User().FindByLogin(u1.Login)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown()

	s := sqlstore.New(db)
	err := s.User().Delete(100)
	assert.NoError(t, err)

	u, _ := s.User().Find(100)
	assert.Nil(t, u)
}

func TestUserRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := models.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)

	newU := &models.User{
		FirstName:   "Test1",
		LastName:    "Test1",
		ThirdName:   "Test1",
		AccessLevel: 2,
		Email:       "dd@yandex.ru",
		Login:       "TestLogin",
	}
	assert.NoError(t, s.User().Update(u.ID, newU))
	assert.Equal(t, newU.FirstName, "Test1")
	assert.Equal(t, newU.LastName, "Test1")
	assert.Equal(t, newU.ThirdName, "Test1")
	assert.Equal(t, newU.AccessLevel, 2)
	assert.Equal(t, newU.Login, "TestLogin")
	assert.Equal(t, newU.Email, "dd@yandex.ru")

	user, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, user.FirstName, "Test1")
	assert.Equal(t, user.LastName, "Test1")
	assert.Equal(t, user.ThirdName, "Test1")
	assert.Equal(t, user.AccessLevel, 2)
	assert.Equal(t, user.Login, "TestLogin")
	assert.Equal(t, user.Email, "dd@yandex.ru")
}
