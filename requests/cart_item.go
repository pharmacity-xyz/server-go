package requests

import "github.com/google/uuid"

type AddCartItem struct {
	UserId    uuid.UUID
	ProductId uuid.UUID
	Quantity  int
}

type UpdateCartItem struct {
	ProductId   uuid.UUID
	ProductName string
	ImageUrl    string
	Price       float64
	Quantity    int64
}
