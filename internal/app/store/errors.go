package store

import "errors"

var (
	// ErrRecordNotFound ...
	ErrRecordNotFound = errors.New("record not found")
	// ErrIncorectEmailOrPassword ...
	ErrIncorectEmailOrPassword = errors.New("Неправильный логин или пароль")
	// ErrNotAuthorized ...
	ErrNotAuthorized = errors.New("Вы не авторизованы")
)
