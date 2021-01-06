package entity_test

import (
	"testing"

	"github.com/task4233/techtrain-mission/gameapi/domain/entity"
)

func TestNewUser(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		arg  string
		want *entity.User
	}{
		{
			name: "正しくユーザを生成できる",
			arg:  "test user",
			want: &entity.User{ID: -1, Name: "test user", Token: ""},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := entity.NewUser(tc.arg)
			if got.ID != tc.want.ID {
				t.Errorf("NewUser().ID = %d, want = %d", got.ID, tc.want.ID)
			}
			if got.Name != tc.want.Name {
				t.Errorf("NewUser().Name = %s, want = %s", got.Name, tc.want.Name)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		arg  *entity.User
		want string
	}{
		{
			name: "正しいユーザ",
			arg:  &entity.User{ID: 0, Name: "test user", Token: "test token"},
			want: "",
		},
		{
			name: "IDが負",
			arg:  &entity.User{ID: -1, Name: "test user", Token: "test token"},
			want: "user.ID is not assigned",
		},
		{
			name: "Nameが無い",
			arg:  &entity.User{ID: 0, Name: "", Token: "test token"},
			want: "user.Name is empty",
		},
		{
			name: "Tokenが無い",
			arg:  &entity.User{ID: 0, Name: "test user", Token: ""},
			want: "user.Token is empty",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.arg.IsValid()
			if err != nil && err.Error() != tc.want {
				t.Errorf("NewUser().Name = %v, want = %v", err.Error(), tc.want)
			}
		})
	}
}
