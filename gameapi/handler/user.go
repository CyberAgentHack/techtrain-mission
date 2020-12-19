package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/task4233/techtrain-mission/gameapi/usecase"
)

// User is struct
type User struct {
	userUC *usecase.User
}

// NewUser returns a pointer for User
func NewUser(userUC *usecase.User) *User {
	return &User{userUC: userUC}
}

// Create handles for POST /user/create
// Create creates user information
// in: name(string), request body required
// out: token(string) StatusOK
func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	req, err := mapUserCreateRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed mapping: %v", err), http.StatusInternalServerError)
		return
	}
	userE, err := u.userUC.CreateWithName(req.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed mapping: %v", err), http.StatusInternalServerError)
		return
	}
	var res UserCreateResponse
	res.Token = userE.Token
	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed marshal: %v", err), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(resJSON)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed marshal: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func mapUserCreateRequest(r *http.Request) (*UserCreateRequest, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed ReadAll: %w", err)
	}
	var req UserCreateRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("failed Unmarshal: %w", err)
	}
	return &req, nil
}

// UserCreateRequest is struct for request on creating user
type UserCreateRequest struct {
	Name string `json:"name"`
}

// UserCreateResponse is struct for response on creating user
type UserCreateResponse struct {
	Token string `json:"token"`
}

// UserGetResponse is struct for response on getting user
type UserGetResponse struct {
	Name string `json:"name"`
}

// UserUpdateRequest is struct for request on updating user
type UserUpdateRequest struct {
	Name string `json:"name"`
}
