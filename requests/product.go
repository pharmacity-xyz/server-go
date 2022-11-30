package requests

import "github.com/google/uuid"

type AddProduct struct {
	ProductName        string
	ProductDescription string
	ImageURL           string
	Stock              int
	Price              int
	CategoryId         uuid.UUID
}

type UpdateProduct struct {
	ProductId          uuid.UUID
	ProductName        string
	ProductDescription string
	ImageURL           string
	Stock              int
	Price              int
	Featured           bool
	CategoryId         uuid.UUID
}
