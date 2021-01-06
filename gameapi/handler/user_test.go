package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/task4233/techtrain-mission/gameapi/handler"
	"github.com/task4233/techtrain-mission/gameapi/log"
	"github.com/task4233/techtrain-mission/gameapi/mock"
	"github.com/task4233/techtrain-mission/gameapi/usecase"
)

func TestUser(t *testing.T) {
	t.Parallel()
	userRepo := mock.NewTestUserRepository()
	userUC := usecase.NewUser(userRepo)
	user := handler.NewUser(userUC)

	TUserCreate(t, user)
	TUserGet(t, user)
	TUserUpdate(t, user)
}

func TUserCreate(t *testing.T, user *handler.User) {
	t.Helper()
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

func TUserGet(t *testing.T, user *handler.User) {
	t.Helper()
	logger := log.MyLogger

	// prepare test data
	reqStruct := handler.UserCreateRequest{Name: "test user"}
	reqBody, err := json.Marshal(reqStruct)
	if err != nil {
		t.Fatalf("failed json.Marshal: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/user/create", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	// Create test user
	user.Create(w, req)
	// Check
	g := w.Result()
	b, err := ioutil.ReadAll(g.Body)
	if err != nil {
		t.Errorf("failed ioutil.ReadAll: %v", err)
	}
	var userCraeteStruct handler.UserCreateResponse
	if err := json.Unmarshal(b, &userCraeteStruct); err != nil {
		t.Fatalf("failed json.Unmarshal: %v", err)
	}
	var token string = userCraeteStruct.Token

	const endpoint = "http://localhost:8080/user/get"

	cases := []struct {
		name         string
		reqMethod    string
		reqHeader    string
		wantStatus   int
		wantResponse string
	}{
		{name: "正常なリクエスト", reqMethod: http.MethodGet, reqHeader: token, wantStatus: http.StatusOK, wantResponse: "test user"},
		{name: "空のリクエストボディ", reqMethod: http.MethodGet, reqHeader: "", wantStatus: http.StatusBadRequest, wantResponse: "x-token must not be empty"},
		{name: "存在しないtoken", reqMethod: http.MethodGet, reqHeader: "test token", wantStatus: http.StatusInternalServerError, wantResponse: "failed GetWithToken: failed userRepo.GetWithToken: failed to Get: No records (&{-1  test token})"},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(tc.reqMethod, endpoint, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-token", tc.reqHeader)

		w := httptest.NewRecorder()

		user.Get(w, req)
		got := w.Result()

		if got.StatusCode != tc.wantStatus {
			t.Errorf("got: %d, wanted: %d", got.StatusCode, tc.wantStatus)
		}

		body, err := ioutil.ReadAll(got.Body)
		if err != nil {
			t.Errorf("failed ioutil.ReadAll: %v", err)
		}
		logger.Debugf("response body: %s\n", string(body))
		// expected want response
		if strings.TrimSpace(string(body)) == tc.wantResponse {
			continue
		}
		var userGetResponse handler.UserGetResponse
		if err := json.Unmarshal(body, &userGetResponse); err != nil {
			t.Fatalf("failed json.Unmarshal: %v", err)
		}
		// check it not empty as token is dynamically generated
		if userGetResponse.Name != tc.wantResponse {
			t.Errorf("got: %s, wanted: %s", userGetResponse.Name, tc.wantResponse)
		}

	}
}

func TUserUpdate(t *testing.T, user *handler.User) {
	t.Helper()
	// prepare test data
	reqStruct := handler.UserCreateRequest{Name: "test user"}
	reqBody, err := json.Marshal(reqStruct)
	if err != nil {
		t.Fatalf("failed json.Marshal: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/user/create", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	// Create test user
	user.Create(w, req)
	// Check
	g := w.Result()
	b, err := ioutil.ReadAll(g.Body)
	if err != nil {
		t.Errorf("failed ioutil.ReadAll: %v", err)
	}
	var userCraeteStruct handler.UserCreateResponse
	if err := json.Unmarshal(b, &userCraeteStruct); err != nil {
		t.Fatalf("failed json.Unmarshal: %v", err)
	}
	var token string = userCraeteStruct.Token

	const endpoint = "http://localhost:8080/user/get"

	cases := []struct {
		name       string
		reqMethod  string
		reqHeader  string
		reqText    string
		wantStatus int
	}{
		{name: "正常なリクエスト", reqMethod: http.MethodPut, reqHeader: token, reqText: "new test user", wantStatus: http.StatusOK},
		{name: "空のリクエストヘッダ", reqMethod: http.MethodPut, reqHeader: "", reqText: "new test user", wantStatus: http.StatusBadRequest},
		{name: "空のリクエストボディ", reqMethod: http.MethodPut, reqHeader: token, reqText: "", wantStatus: http.StatusBadRequest},
		{name: "存在しないtoken", reqMethod: http.MethodPut, reqHeader: "test token", reqText: "new test user", wantStatus: http.StatusInternalServerError},
	}

	for _, tc := range cases {
		reqStruct := handler.UserUpdateRequest{Name: tc.reqText}
		reqBody, err := json.Marshal(reqStruct)
		if err != nil {
			t.Fatalf("failed json.Marshal: %v", err)
		}
		req := httptest.NewRequest(tc.reqMethod, endpoint, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-token", tc.reqHeader)

		w := httptest.NewRecorder()

		user.Update(w, req)
		got := w.Result()

		if got.StatusCode != tc.wantStatus {
			t.Errorf("got: %d, wanted: %d", got.StatusCode, tc.wantStatus)
		}
	}

}
