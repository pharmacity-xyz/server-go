package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/responses"
	"github.com/pharmacity-xyz/server-go/types"
)

type Orders struct {
	OrderService *models.OrderService
}

func (o Orders) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.ServiceResponse[*[]responses.OrderOverviewResponse]{
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

	products, err := o.OrderService.GetOrders(userId)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = products
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (o Orders) GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.ServiceResponse[*responses.OrderDetailsResponse]{
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

	orderId := chi.URLParam(r, "orderId")
	products, err := o.OrderService.GetOrderDetails(userId, orderId)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = products
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (o Orders) GetOrdersForAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.ServiceResponse[*[]responses.OrderOverviewResponse]{
		Message: "",
	}

	err := AuthorizeAdmin(r)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	product, err := o.OrderService.GetOrdersForAdmin()
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = product
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (o Orders) GetOrdersPerMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.ServiceResponse[[]uint]{
		Message: "",
	}

	err := AuthorizeAdmin(r)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	q := r.URL.Query()
	yearStr := q.Get("year")
	monthStr := q.Get("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	product, err := o.OrderService.GetOrderPerMonth(year, month)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = product
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
