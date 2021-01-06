package infra

import (
	"os"
	"testing"

	"github.com/task4233/techtrain-mission/gameapi/domain/entity"
	"github.com/task4233/techtrain-mission/gameapi/mock"
	"github.com/task4233/tecktrain-mission/gameapi/log"
)

var (
	logger = log.MyLogger
)

// This test does not use mock
func TestUserInfraWithMySQL(t *testing.T) {
	// scenario
	// 1. Store User with name
	// 2. Get User information
	// 3. Update User information

	db, err := NewDB()
	if err != nil {
		t.Errorf("failed NewDB: %v", err)
		os.Exit(1)
	}
	if db == nil {
		t.Errorf("failed to connect NewDB")
		os.Exit(1)
	}
	defer func() {
		cerr := db.Close()
		if err != nil {
			logger.Debugf("failed Close(): %w", cerr)
		}
	}()
	userRepo := NewUserRepository(db)

	// 1. Store
	userE := entity.NewUser("test user")
	if err := userRepo.Store(userE); err != nil {
		t.Errorf("failed Store: %v", err)
	}

	// 2. Get
	var recordE entity.User
	recordE.Token = userE.Token
	if err := userRepo.Get(&recordE); err != nil {
		t.Errorf("failed Get: %v", err)
	}
	if err := recordE.IsValid(); err != nil {
		t.Errorf("invalid user entity: %v, %v", recordE, err)
	}

	// 3. Update
	recordE.Name = "test user 2"
	if err := userRepo.Update(&recordE); err != nil {
		t.Errorf("failed Update: %v", err)
	}
	if err := recordE.IsValid(); err != nil {
		t.Errorf("invalid user entity: %v, %v", recordE, err)
	}
	var testE *entity.User = entity.NewUser("")
	testE.Token = recordE.Token
	if err := userRepo.Get(testE); err != nil {
		t.Errorf("failed Get: %v", err)
	}
	if err := testE.IsValid(); err != nil {
		t.Errorf("invalid user entity: %v, %v", testE, err)
	}
	if testE.Token != recordE.Token {
		t.Errorf("might fail to update, actual: %v, wanted: %v", testE, recordE)
	}
}

func TestUserInfraWithMock(t *testing.T) {
	// scenario
	// 1. Store User with name
	// 2. Get User information
	// 3. Update User information

	userRepo := mock.NewTestUserRepository()

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
