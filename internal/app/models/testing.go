package models

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Login:    "Hello",
		Email:    "testing@yandex.ru",
		Password: "MyPassword",
	}
}
