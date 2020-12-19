package infra

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/task4233/techtrain-mission/gameapi/domain/entity"
	"github.com/task4233/techtrain-mission/gameapi/repository"
)

// UserRepository implements repository.User
type UserRepository struct {
	db *sqlx.DB
}

var _ repository.User = &UserRepository{}

// NewUserRepository returns the pointer for UserRepository
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Store stores user entity
func (u *UserRepository) Store(user *entity.User) error {
	dto := newDTO(user)

	if _, err := u.db.NamedExec(`INSERT INTO users (name, token) VALUES (:name, :token)`, dto); err != nil {
		return fmt.Errorf("failed db.NamedExec: %w", err)
	}

	return nil
}

// Get gets user entity
func (u *UserRepository) Get(user *entity.User) error {
	dto := userDTO{}

	if err := u.db.Get(&dto, `SELECT * FROM users WHERE token=?`, user.Token); err != nil {
		return fmt.Errorf("failed db.Select: %w", err)
	}
	user.ID = dto.ID
	user.Name = dto.Name
	return nil
}

// Update updates user entity
func (u *UserRepository) Update(user *entity.User) error {
	var dto *userDTO = newDTO(user)

	if _, err := u.db.NamedExec(`UPDATE users SET name=? WHERE token=?`, dto); err != nil {
		return fmt.Errorf("failed db.NamedExec: %w", err)
	}
	return nil
}

func newDTO(user *entity.User) *userDTO {
	return &userDTO{
		ID:    user.ID,
		Name:  user.Name,
		Token: user.Token,
	}
}

type userDTO struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Token string `db:"token"`
}
