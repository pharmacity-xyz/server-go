package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/requests"
	"github.com/pharmacity-xyz/server-go/responses"
)

type CartItems struct {
	CartItemService *models.CartItemService
}

func (c CartItems) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request requests.AddCartItem
	var response = responses.CategoryResponse[*models.CartItem]{
		Message: "",
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusBadRequest)
		return
	}

	newCartItem := models.CartItem{
		UserId:    request.UserId,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
	}
	cartItem, err := c.CartItemService.Add(&newCartItem)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = cartItem
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
