package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pharmacity-xyz/server-go/models"
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
