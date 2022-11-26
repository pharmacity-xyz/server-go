package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pharmacity-xyz/server-go/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response = utils.HealthCheckResponse{
		Data:    "",
		Success: true,
	}
	json.NewEncoder(w).Encode(response)
}
