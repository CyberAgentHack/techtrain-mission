package repository

import "github.com/task4233/techtrain-mission/gameapi/domain/entity"

// User is for persistence and reconstruction for user
type User interface {
	Store(user *entity.User) error
	Get(user *entity.User) error
	Update(user *entity.User) error
}
