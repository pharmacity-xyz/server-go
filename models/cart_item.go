package models

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type CartItemService struct {
	DB *sql.DB
}

type CartItem struct {
	UserId    uuid.UUID
	ProductId uuid.UUID
	Quantity  int
}

func (cis CartItemService) Add(newCartItem *CartItem) (*CartItem, error) {
	_, err := cis.DB.Exec(`
		INSERT INTO cart_items (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
	`, newCartItem.UserId, newCartItem.ProductId, newCartItem.Quantity)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	return newCartItem, nil
}
