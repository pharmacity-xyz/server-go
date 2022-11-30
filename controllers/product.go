package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

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

func (p Products) GetProductByCategoryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.ServiceResponse[[]*models.Product]{
		Message: "",
	}

	categoryId := chi.URLParam(r, "categoryId")
	products, err := p.ProductService.GetProductByCategoryId(categoryId)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = products
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (p Products) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.ServiceResponse[[]*models.Product]{
		Message: "",
	}

	searchWord := chi.URLParam(r, "searchWord")
	lowerSearchWord := strings.ToLower(searchWord)
	products, err := p.ProductService.Search(lowerSearchWord)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = products
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (p Products) FeaturedProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = types.ServiceResponse[[]*models.Product]{
		Message: "",
	}

	products, err := p.ProductService.FeaturedProducts()
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = products
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (p Products) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request requests.UpdateProduct
	var response = responses.CategoryResponse[*models.Product]{
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
		ProductId:          request.ProductId,
		ProductName:        request.ProductName,
		ProductDescription: request.ProductDescription,
		ImageURL:           request.ImageURL,
		Stock:              request.Stock,
		Price:              request.Price,
		Featured:           request.Featured,
		CategoryId:         request.CategoryId,
	}
	product, err := p.ProductService.Update(&newProduct)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = product
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (p Products) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = responses.CategoryResponse[bool]{
		Message: "",
	}
	productId := chi.URLParam(r, "productId")

	err := AuthorizeAdmin(r)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	err = p.ProductService.Delete(productId)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	response.Data = true
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
