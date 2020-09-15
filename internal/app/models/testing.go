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

// TestManufacturer ...
func TestManufacturer(t *testing.T) *Manufacturer {
	m := &Manufacturer{
		Name:        "Тестовый производитель",
		Description: "Каширка 19к2",
	}
	m.Units = append(m.Units, &Unit{
		ID:      1,
		Name:    "test",
		Address: "Net",
	})
	return m
}
