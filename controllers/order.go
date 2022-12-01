package controllers

import (
	"encoding/json"
	"net/http"

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
