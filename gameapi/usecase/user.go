package usecase

import (
	"fmt"

	"github.com/task4233/techtrain-mission/gameapi/domain/entity"
	"github.com/task4233/techtrain-mission/gameapi/repository"
)

// User manages the UserUsecase
type User struct {
	userRepo repository.User
}

// NewUser return the pointer for usecase.User
func NewUser(userRepo repository.User) *User {
	return &User{userRepo: userRepo}
}

// CreateWithName creates user information
// Simultaneously, user.ID and user.Token are assigned and stored
func (u *User) CreateWithName(name string) (*entity.User, error) {
	userE := entity.NewUser(name)
	if err := u.userRepo.Store(userE); err != nil {
		return nil, fmt.Errorf("failed userRepo.Store: %w", err)
	}
	return userE, nil
}

// GetWithToken gets user information with token
func (u *User) GetWithToken(token string) (*entity.User, error) {
	userE := entity.NewUser("")
	userE.Token = token
	if err := u.userRepo.Get(userE); err != nil {
		return nil, fmt.Errorf("failed userRepo.GetWithToken: %w", err)
	}
	return userE, nil
}

// UpdateWithToken updates user information with token
func (u *User) UpdateWithToken(token string, name string) error {
	userE := entity.NewUser(name)
	userE.Token = token
	if err := u.userRepo.Update(userE); err != nil {
		return fmt.Errorf("failed userRepo.Update: %w", err)
	}
	return nil
}
