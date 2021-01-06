package config_test

import (
	"testing"

	"github.com/task4233/techtrain-mission/gameapi/config"
)

func TestPort(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		want string
	}{
		{
			name: "正しくポートを取得できる",
			want: "8080",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := config.Port(); got != tc.want {
				t.Errorf("Port() = %s, want = %s", got, tc.want)
			}
		})
	}
}

func TestDSN(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		want string
	}{
		{
			name: "正しくDSNを取得できる",
			want: "game:gamepass@tcp(127.0.0.1:3306)/game?parseTime=true&collation=utf8mb4_bin",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := config.DSN(); got != tc.want {
				t.Errorf("DSN() = %s, want = %s", got, tc.want)
			}
		})
	}
}

func TestIsDev(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		want bool
	}{
		{
			name: "正しくDEVを取得できる",
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := config.IsDev(); got != tc.want {
				t.Errorf("IsDev() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestIsTest(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		want bool
	}{
		{
			name: "正しくDEVを取得できる",
			want: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := config.IsTest(); got != tc.want {
				t.Errorf("IsTest() = %v, want = %v", got, tc.want)
			}
		})
	}
}
