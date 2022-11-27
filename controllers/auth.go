package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/utils"
)

type Users struct {
	UserService *models.UserService
}

func (u Users) Register(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	var response = utils.RegisterResponse{
		Data:    "",
		Message: "",
	}

	w.Header().Set("Content-Type", "application/json")
	user := models.User{}
	user1, err := u.UserService.Register(&user, password)
	if err != nil {
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
	}
	response.Data = user1.UserId.String()

	json.NewEncoder(w).Encode(response)
}

func (u Users) Login(w http.ResponseWriter, r *http.Request) {
	var response = utils.LoginResponse{
		Data:    "",
		Message: "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (u Users) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var response = utils.LoginResponse{
		Data:    "",
		Message: "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (u Users) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var response = utils.LoginResponse{
		Data:    "",
		Message: "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
