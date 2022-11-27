package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/requests"
	"github.com/pharmacity-xyz/server-go/responses"
)

type Users struct {
	UserService *models.UserService
}

func (u Users) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request requests.Register
	var response = responses.RegisterResponse{
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
		Email:       request.Email,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		City:        request.City,
		Country:     request.Country,
		CompanyName: request.CompanyName,
	}
	_, err = u.UserService.Register(&user, request.Password)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = user.UserId.String()
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (u Users) Login(w http.ResponseWriter, r *http.Request) {
	var response = responses.LoginResponse{
		Data:    "",
		Message: "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (u Users) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var response = responses.LoginResponse{
		Data:    "",
		Message: "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (u Users) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var response = responses.LoginResponse{
		Data:    "",
		Message: "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
