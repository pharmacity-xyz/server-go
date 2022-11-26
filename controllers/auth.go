package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pharmacity-xyz/server-go/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = utils.RegisterResponse{
		Data:    "",
		Success: true,
	}
	json.NewEncoder(w).Encode(response)
}
