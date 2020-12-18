package entity

import "github.com/google/uuid"

// User indicates user infomation
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
		ID:    0,
		Name:  name,
		Token: uuid.New().String(),
	}
}
