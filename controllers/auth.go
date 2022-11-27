package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/requests"
	"github.com/pharmacity-xyz/server-go/responses"
)

type Users struct {
	UserService *models.UserService
}

type Claims struct {
	UserId uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
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
		Email:       strings.ToLower(request.Email),
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
	w.Header().Set("Content-Type", "application/json")
	var request requests.Login
	var response = responses.LoginResponse{
		Data:    "",
		Message: "",
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusBadRequest)
		return
	}

	user, err := u.UserService.Login(request.Email, request.Password)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserId: user.UserId,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
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
