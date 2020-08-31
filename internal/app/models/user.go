package models

import (
	"crypto/sha512"
	"encoding/base64"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User implements user
type User struct {
	ID                int    `json:"id,omitempty"`
	Login             string `json:"login,omitempty"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
	Email             string `json:"email,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	ThirdName         string `json:"third_name,omitempty"`
	AccessLevel       int    `json:"access_level,omitempty"`
	LastLogin         int    `json:"last_login,omitempty"`
	CreatedAt         int    `json:"created_at,omitempty"`
}

// Validate validates User structure
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Login, validation.Required),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

// Sanitize deletes all private data before response
func (u *User) Sanitize() {
	u.Password = ""
}

// BeforeCreate encryptes password for User
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	u.AccessLevel = 1
	return nil
}

// ComparePassword checks is given password is correct
func (u *User) ComparePassword(password string) bool {
	enc, err := encryptString(password)
	if err != nil {
		return false
	}
	return u.EncryptedPassword == enc
}

func encryptString(password string) (string, error) {
	hash := []byte(password)
	dig := sha512.Sum512(hash)
	for i := 1; i < 5000; i++ {
		dig = sha512.Sum512(append(dig[:], hash[:]...))
	}
	result := base64.StdEncoding.EncodeToString(dig[:])
	return result, nil
}

/*
 ** Checks for USER
 */

// IsManager check is user has manager access level
func (u *User) IsManager() bool {
	return u.AccessLevel > 1
}

// IsDirector checks is user has Director access level
func (u *User) IsDirector() bool {
	return u.AccessLevel > 2
}

// IsRRS checks is user has RRS access level
func (u *User) IsRRS() bool {
	return u.AccessLevel > 4
}

// IsAdmin checks is user has admin access level
func (u *User) IsAdmin() bool {
	return u.AccessLevel > 3
}
