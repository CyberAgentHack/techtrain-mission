package entity

import (
	"errors"

	"github.com/google/uuid"
)

// User indicates user information
type User struct {
	// ID is assigned on calling /user/create
	ID int
	// Name is defined by user in request body on calling /user/create
	Name string
	// Token is generated on calling /user/create
	Token string
}

// NewUser returns the pointer for making new
func NewUser(name string) *User {
	return &User{
		ID:    -1, // this is invalid ID
		Name:  name,
		Token: uuid.New().String(),
	}
}

// IsValid validates user entity
// It might be moved domain service
func (u *User) IsValid() error {
	if u.ID == -1 {
		return errors.New("user.ID is not assigned")
	}
	if len(u.Name) == 0 {
		return errors.New("user.Name is empty")
	}
	if len(u.Token) == 0 {
		return errors.New("user.Token is empty")
	}
	return nil
}
