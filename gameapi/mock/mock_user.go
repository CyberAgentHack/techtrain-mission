package mock

import (
	"fmt"

	"github.com/task4233/techtrain-mission/gameapi/domain/entity"
	"github.com/task4233/techtrain-mission/gameapi/repository"
)

// TestUserRepository implements repository.User
type TestUserRepository struct {
	records [](*entity.User)
}

var _ repository.User = &TestUserRepository{}

// NewTestUserRepository returns the pointer for UserRepository
func NewTestUserRepository() *TestUserRepository {
	return &TestUserRepository{
		records: [](*entity.User){},
	}
}

// Store stores user entity
func (t *TestUserRepository) Store(user *entity.User) error {
	// user.ID is AUTO_INCREMENT
	user.ID = len(t.records)
	userE := user
	t.records = append(t.records, userE)
	return nil
}

// Get gets user entity
func (t *TestUserRepository) Get(user *entity.User) error {
	for _, record := range t.records {
		if record.Token == user.Token {
			user.ID = record.ID
			user.Name = record.Name
			return nil
		}
	}
	return fmt.Errorf("failed to Get: No records (%v)", user)
}

// Update updates user entity
func (t *TestUserRepository) Update(user *entity.User) error {
	for _, record := range t.records {
		if record.Token == user.Token {
			user.ID = record.ID
			record.Name = user.Name
			return nil
		}
	}
	return fmt.Errorf("failed to Update: No records (%v)", user)
}
