package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/requests"
	"github.com/pharmacity-xyz/server-go/responses"
)

type Auths struct {
	UserService *models.UserService
}

func (a Auths) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request requests.Register
	var response = responses.AuthResponse[string]{
		Data:    "",
		Message: "",
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusBadRequest)
		return
	}

	user := models.User{
		UserId:      uuid.New(),
		Email:       strings.ToLower(request.Email),
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		City:        request.City,
		Country:     request.Country,
		CompanyName: request.CompanyName,
	}
	_, err = a.UserService.Register(&user, request.Password)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = user.UserId.String()
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (a Auths) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request requests.Login
	var response = responses.AuthResponse[string]{
		Data:    "",
		Message: "",
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusBadRequest)
		return
	}

	user, err := a.UserService.Login(request.Email, request.Password)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	fmt.Printf("User Role: %s\n", user.Role)
	tokenString, expirationTime, err := CreateJWT(user.UserId, user.Role)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	setCookie(w, COOKIE_TOKEN, tokenString, expirationTime)
	response.Data = tokenString
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (a Auths) ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newPassword := r.FormValue("new_password")
	var response = responses.AuthResponse[string]{
		Data:    "",
		Message: "",
	}

	token, err := readCookie(r, COOKIE_TOKEN)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	userId, _, err := ParseJWT(token)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	success, err := a.UserService.ChangePassword(userId, newPassword)
	if !success || err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (a Auths) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var response = responses.AuthResponse[string]{
		Data:    "",
		Message: "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
