package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/requests"
	"github.com/pharmacity-xyz/server-go/responses"
	"github.com/pharmacity-xyz/server-go/types"
)

type Users struct {
	UserService *models.UserService
}

func (u Users) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.ServiceResponse[[]*models.User]{
		Message: "",
	}

	err := AuthorizeAdmin(r)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	users, err := u.UserService.GetAll()
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = users
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (u Users) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request requests.UpdateUser
	var response = types.ServiceResponse[*models.User]{
		Message: "",
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusBadRequest)
		return
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

	newUser := models.User{
		UserId:      userId,
		Email:       request.Email,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		City:        request.City,
		Country:     request.Country,
		CompanyName: request.CompanyName,
	}
	err = u.UserService.Update(&newUser)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = &newUser
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
