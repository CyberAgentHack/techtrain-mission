package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/task4233/techtrain-mission/gameapi/handler"
	"github.com/task4233/techtrain-mission/gameapi/infra"
	"github.com/task4233/techtrain-mission/gameapi/usecase"
)

func TestUserCreate(t *testing.T) {
	t.Parallel()

	const endpoint = "http://localhost:8080/user/create"

	cases := []struct {
		name       string
		reqMethod  string
		reqText    string
		wantStatus int
	}{
		{name: "正常なリクエスト", reqMethod: http.MethodPost, reqText: "test user", wantStatus: http.StatusOK},
		{name: "空のリクエストボディ", reqMethod: http.MethodPost, reqText: "", wantStatus: http.StatusBadRequest},
	}

	userRepo := infra.NewTestUserRepository()
	userUC := usecase.NewUser(userRepo)
	user := handler.NewUser(userUC)

	for _, tc := range cases {
		reqStruct := handler.UserCreateRequest{Name: tc.reqText}
		reqBody, err := json.Marshal(reqStruct)
		if err != nil {
			t.Fatalf("failed json.Marshal: %v", err)
		}
		req := httptest.NewRequest(tc.reqMethod, endpoint, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		user.Create(w, req)
		got := w.Result()

		if got.StatusCode != tc.wantStatus {
			t.Errorf("got: %d, wanted: %d", got.StatusCode, tc.wantStatus)
		}

		body, err := ioutil.ReadAll(got.Body)
		if err != nil {
			t.Errorf("failed ioutil.ReadAll: %v", err)
		}
		// check it not empty as token is dynamically generated
		if len(string(body)) == 0 {
			t.Error("response body is empty")
		}
	}
}
