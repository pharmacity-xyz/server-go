package responses

type HealthCheckResponse struct {
	Data    string `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
