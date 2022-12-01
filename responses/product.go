package responses

import "github.com/google/uuid"

type ProductResponse[T any] struct {
	Data    T      `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type OrderDetailsProductResponse struct {
	ProductId   uuid.UUID
	ProductName string
	ImageUrl    string
	Quantity    int64
	TotalPrice  float64
}
