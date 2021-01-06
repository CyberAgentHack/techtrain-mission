package usecase_test

import (
	"testing"

	"github.com/task4233/techtrain-mission/gameapi/mock"
	"github.com/task4233/techtrain-mission/gameapi/usecase"
)

func TestCreateWithName(t *testing.T) {
	t.Parallel()
	userRepo := mock.NewTestUserRepository()
	userUC := usecase.NewUser(userRepo)

	cases := []struct {
		name    string
		arg     string
		wantErr string
	}{
		{
			name:    "正しく保存できる",
			arg:     "test user",
			wantErr: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := userUC.CreateWithName(tc.arg); err != nil && err.Error() != tc.wantErr {
				t.Errorf("userUC.CreateWithName() error = %s, want = %s", err.Error(), tc.wantErr)
			}
		})
	}
}

func TestGetWithToken(t *testing.T) {
	t.Parallel()
	userRepo := mock.NewTestUserRepository()
	userUC := usecase.NewUser(userRepo)

	cases := []struct {
		name    string
		wantErr string
	}{
		{
			name:    "正しく取得できる",
			wantErr: "",
		},
	}

	var arg string = "test user"
	userE, err := userUC.CreateWithName(arg)
	if err != nil {
		t.Fatalf("failed CreateWithName(): %v", err)
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			userE, err := userUC.GetWithToken(userE.Token)
			if userE.Name != arg {
				t.Errorf("userUC.GetWithToken() = %s, want = %s", userE.Name, arg)
			}
			if err != nil && err.Error() != tc.wantErr {
				t.Errorf("userUC.GetWithToken() error = %s, want = %s", err.Error(), tc.wantErr)
			}
		})
	}
}

func TestUpdateWithToken(t *testing.T) {
	t.Parallel()
	userRepo := mock.NewTestUserRepository()
	userUC := usecase.NewUser(userRepo)

	cases := []struct {
		name    string
		arg     string
		wantErr string
	}{
		{
			name:    "正しく更新できる",
			arg:     "test user 2",
			wantErr: "",
		},
	}

	var arg string = "test user"
	userE, err := userUC.CreateWithName(arg)
	if err != nil {
		t.Fatalf("failed CreateWithName(): %v", err)
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := userUC.UpdateWithToken(userE.Token, tc.arg); err != nil && err.Error() != tc.wantErr {
				t.Errorf("userUC.UpdateWithToken() error = %s, want = %s", err.Error(), tc.wantErr)
			}

			userE2, err := userUC.GetWithToken(userE.Token)
			if userE2.Name != tc.arg {
				t.Errorf("userUC.UpdateWithToken() = %s, want = %s", userE2.Name, tc.arg)
			}
			if err != nil && err.Error() != tc.wantErr {
				t.Errorf("userUC.GetWithToken() error = %s, want = %s", err.Error(), tc.wantErr)
			}
		})
	}
}
