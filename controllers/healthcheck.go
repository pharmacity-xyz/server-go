package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pharmacity-xyz/server-go/responses"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = responses.HealthCheckResponse{
		Data:    "",
		Success: true,
	}
	json.NewEncoder(w).Encode(response)
}
