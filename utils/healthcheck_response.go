package utils

type HealthCheckResponse struct {
	Data    string `json:"data"`
	Success bool   `json:"type"`
	Message string `json:"message"`
}
