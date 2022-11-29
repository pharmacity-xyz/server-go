package requests

import "github.com/google/uuid"

type AddCategory struct {
	CategoryName string
}

type UpdateCategory struct {
	CategoryId uuid.UUID
	Name       string
}
