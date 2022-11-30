package requests

import "github.com/google/uuid"

type AddCartItem struct {
	UserId    uuid.UUID
	ProductId uuid.UUID
	Quantity  int
}
