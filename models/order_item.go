package models

import "github.com/google/uuid"

type OrderItem struct {
	OrderId    uuid.UUID
	ProductId  uuid.UUID
	Quantity   int64
	TotalPrice float64
}
