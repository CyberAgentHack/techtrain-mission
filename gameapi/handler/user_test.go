package handler_test

import (
	"net/http"
	"testing"
)

func TestUserCreate(t *testing.T) {
	t.Parallel()

	const endpoint = "http://localhost:8080/user/create"

	cases := []struct {
		name         string
		reqMethod    string
		reqText      string
		wantStatus   int
		wantResponse string
	}{
		{name: "正常なリクエスト", reqMethod: http.MethodPost, reqText: "test user", wantStatus: http.StatusOK, wantResponse: ""},
	}
}
