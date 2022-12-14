package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pharmacity-xyz/server-go/models"
	"github.com/pharmacity-xyz/server-go/requests"
	"github.com/pharmacity-xyz/server-go/responses"
)

type Categories struct {
	CategoryService *models.CategoryService
}

func (c Categories) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = responses.CategoryResponse[[]*models.Category]{
		Message: "",
	}

	categories, err := c.CategoryService.GetAll()
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = categories
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (c Categories) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request requests.AddCategory
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

func (c Categories) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request requests.UpdateCategory
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
		CategoryId: request.CategoryId,
		Name:       request.Name,
	}
	category, err := c.CategoryService.Update(&newCategory)
	if err != nil {
		response.Message = err.Error()
		responses.JSONError(w, response, http.StatusInternalServerError)
		return
	}

	response.Data = category
	response.Success = true
	json.NewEncoder(w).Encode(response)
}

func (c Categories) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = responses.CategoryResponse[bool]{
		Message: "",
	}
	categoryId := chi.URLParam(r, "categoryId")

	err := AuthorizeAdmin(r)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		responses.JSONError(w, response, http.StatusUnauthorized)
		return
	}

	err = c.CategoryService.Delete(categoryId)
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
