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

// TestUnit ...
func TestUnit(t *testing.T) *Unit {
	t.Helper()

	return &Unit{
		Name:    "Test Unit",
		Address: "Test Address",
	}
}
