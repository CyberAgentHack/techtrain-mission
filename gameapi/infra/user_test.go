package infra

import (
	"fmt"
	"testing"

	"github.com/task4233/techtrain-mission/gameapi/domain/entity"
	"github.com/task4233/techtrain-mission/gameapi/repository"
)

func TestUserInfra(t *testing.T) {
	// scenario
	// 1. Store User with name
	// 2. Get User information
	// 3. Update User information

	userRepo := NewTestUserRepository()

	// 1. Store
	userE := entity.NewUser("test user")
	if err := userRepo.Store(userE); err != nil {
		t.Errorf("failed Store: %v", err)
	}
	if err := userE.IsValid(); err != nil {
		t.Errorf("invalid user entity: %w", err)
	}

	// 2. Get
	var recordE entity.User
	recordE.Token = userE.Token
	if err := userRepo.Get(&recordE); err != nil {
		t.Errorf("failed Get: %v", err)
	}
	if err := recordE.IsValid(); err != nil {
		t.Errorf("invalid user entity: %w", err)
	}

	// 3. Update
	recordE.Name = "test user 2"
	if err := userRepo.Update(&recordE); err != nil {
		t.Errorf("failed Update: %v", err)
	}
	if err := recordE.IsValid(); err != nil {
		t.Errorf("invalid user entity: %w", err)
	}
	var testE *entity.User = entity.NewUser("")
	testE.Token = recordE.Token
	if err := userRepo.Get(testE); err != nil {
		t.Errorf("failed Get: %v", err)
	}
	if err := testE.IsValid(); err != nil {
		t.Errorf("invalid user entity: %w", err)
	}
	if testE.Token != recordE.Token {
		t.Errorf("might fail to update, actual: %v, wanted: %v", testE, recordE)
	}
}

type TestUserRepository struct {
	records [](*entity.User)
}

var _ repository.User = &TestUserRepository{}

func NewTestUserRepository() *TestUserRepository {
	return &TestUserRepository{
		records: [](*entity.User){},
	}
}

func (t *TestUserRepository) Store(user *entity.User) error {
	// user.ID is AUTO_INCREMENT
	user.ID = len(t.records)
	userE := user
	t.records = append(t.records, userE)
	return nil
}

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

func (t *TestUserRepository) Update(user *entity.User) error {
	for _, record := range t.records {
		if record.Token == user.Token {
			user.ID = record.ID
			user.Name = record.Name
			return nil
		}
	}
	return fmt.Errorf("failed to Update: No records (%v)", user)
}
