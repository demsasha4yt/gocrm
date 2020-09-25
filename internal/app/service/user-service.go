package service

import (
	"context"

	"github.com/demsasha4yt/gocrm.git/internal/app/models"
)

// UsersServiceInterface ...
type UsersServiceInterface interface {
	Create(context.Context, *models.User) error
	Find(context.Context, int) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByLogin(context.Context, string) (*models.User, error)
	Delete(context.Context, int) error
	Update(context.Context, int, *models.User) error
}

// UsersService ...
type UsersService struct {
	service *Service
}

// Create creates user
func (s *UsersService) Create(ctx context.Context, u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return s.service.store.Users().Create(ctx, u)
}

// Find user
func (s *UsersService) Find(ctx context.Context, id int) (*models.User, error) {
	return s.service.store.Users().Find(ctx, id)
}

// FindByEmail user by Email
func (s *UsersService) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.service.store.Users().FindByEmail(ctx, email)
}

// FindByLogin user by Login
func (s *UsersService) FindByLogin(ctx context.Context, login string) (*models.User, error) {
	return s.service.store.Users().FindByLogin(ctx, login)
}

// Delete deletes user
func (s *UsersService) Delete(ctx context.Context, id int) error {
	return s.service.store.Users().Delete(ctx, id)
}

// Update user
func (s *UsersService) Update(ctx context.Context, id int, u *models.User) error {
	u.EncryptedPassword = "a" // to avoid validate passoword field
	if err := u.Validate(); err != nil {
		return err
	}
	return s.service.store.Users().Update(ctx, id, u)
}
