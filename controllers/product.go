package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/requests"
	"github.com/pharmacity-xyz/server-go/responses"
)

type Products struct {
	ProductService *models.ProductService
}

func (p Products) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request requests.AddProduct
	var response = responses.CategoryResponse[*models.Category]{
		Message: "",
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusBadRequest)
		return
	}

	err = AuthorizeAdmin(r)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	newCategory := models.Category{
		CategoryId: uuid.New(),
		Name:       request.CategoryName,
	}
	categories, err := c.CategoryService.Add(&newCategory)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = categories
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
