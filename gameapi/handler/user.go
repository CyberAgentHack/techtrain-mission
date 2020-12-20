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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed ReadAll: %v", err), http.StatusInternalServerError)
		return
	}
	var req UserCreateRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, fmt.Sprintf("failed Unmarshal: %v", err), http.StatusInternalServerError)
		return
	}
	if len(req.Name) == 0 {
		http.Error(w, "name must not be empty", http.StatusBadRequest)
		return
	}

	userE, err := u.userUC.CreateWithName(req.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed CreateWithName: %v", err), http.StatusInternalServerError)
		return
	}

	res := UserCreateResponse{Token: userE.Token}
	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed Marshal: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resJSON)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed w.Write: %v", err), http.StatusInternalServerError)
		return
	}
}

// Get handles GET /user/get
// Get user information
// in: x-token(string), request header required
// out: name(string) StatusOK
func (u *User) Get(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	if len(token) == 0 {
		http.Error(w, "x-token must not be empty", http.StatusBadRequest)
		return
	}

	userE, err := u.userUC.GetWithToken(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed GetWithToken: %v", err), http.StatusInternalServerError)
		return
	}

	res := UserGetResponse{Name: userE.Name}
	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed Marshal: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(resJSON)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed w.Write: %v", err), http.StatusInternalServerError)
		return
	}
}

// Update handles PUT /user/update
// Update user information
// in: x-token(string), request header required
//     name(string), request body required
// out: StatusOK
func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	if len(token) == 0 {
		http.Error(w, "x-token must not be empty", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed ReadAll: %v", err), http.StatusInternalServerError)
		return
	}
	var req UserUpdateRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, fmt.Sprintf("failed Unmarshal: %v", err), http.StatusInternalServerError)
		return
	}
	if len(req.Name) == 0 {
		http.Error(w, "name must not be empty", http.StatusBadRequest)
		return
	}

	if err := u.userUC.UpdateWithToken(token, req.Name); err != nil {
		http.Error(w, fmt.Sprintf("failed GetWithToken: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
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
