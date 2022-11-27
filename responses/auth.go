package responses

type RegisterResponse struct {
	Data    string `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Data    string `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
