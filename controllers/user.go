package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/responses"
)

type Users struct {
	UserService *models.UserService
}

func (u Users) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = responses.UserResponse[[]*models.User]{
		Message: "",
	}
	token, err := readCookie(r, COOKIE_TOKEN)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	_, role, err := ParseJWT(token)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}
	if role != "Admin" {
		response.Message = "unauthorized"
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	users, err := u.UserService.GetAll()
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	response.Data = users
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
