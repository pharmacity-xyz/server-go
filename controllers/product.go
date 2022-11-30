package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/requests"
	"github.com/pharmacity-xyz/server-go/responses"
	"github.com/pharmacity-xyz/server-go/types"
)

type Products struct {
	ProductService *models.ProductService
}

func (p Products) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request requests.AddProduct
	var response = types.ServiceResponse[*models.Product]{
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

	newProduct := models.Product{
		ProductId:          uuid.New(),
		ProductName:        request.ProductName,
		ProductDescription: request.ProductDescription,
		ImageURL:           request.ImageURL,
		Stock:              request.Stock,
		Price:              request.Price,
		Featured:           false,
		CategoryId:         request.CategoryId,
	}
	product, err := p.ProductService.Add(&newProduct)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = product
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (p Products) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.ServiceResponse[[]*models.Product]{
		Message: "",
	}

	products, err := p.ProductService.GetAll()
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = products
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (p Products) GetProductByProductId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.ServiceResponse[*models.Product]{
		Message: "",
	}

	productId := chi.URLParam(r, "productId")
	product, err := p.ProductService.GetProductByProductId(productId)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = product
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
