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

type CartItemWithProduct struct {
	ProductId   uuid.UUID
	ProductName string
	ImageUrl    string
	Price       float64
	Quantity    int64
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

func (cis CartItemService) GetAll(userId uuid.UUID) ([]*CartItemWithProduct, error) {
	var cartItems []*CartItemWithProduct
	rows, err := cis.DB.Query(`
		SELECT products.product_id, product_name, image_url, price, quantity 
		FROM cart_items 
		JOIN products ON 
		products.product_id = cart_items.product_id
		WHERE user_id = $1
	`, userId)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cartItemWithProduct CartItemWithProduct
		if err := rows.Scan(
			&cartItemWithProduct.ProductId,
			&cartItemWithProduct.ProductName,
			&cartItemWithProduct.ImageUrl,
			&cartItemWithProduct.Price,
			&cartItemWithProduct.Quantity,
		); err != nil {
			return nil, fmt.Errorf("fail: %w", err)
		}
		cartItems = append(cartItems, &cartItemWithProduct)
	}
	return cartItems, nil
}

func (cis *CartItemService) Count(userId uuid.UUID) (int, error) {
	var cartItems []*CartItemWithProduct
	rows, err := cis.DB.Query(`
		SELECT products.product_id 
		FROM cart_items 
		JOIN products ON 
		products.product_id = cart_items.product_id
		WHERE user_id = $1
	`, userId)
	if err != nil {
		return 0, fmt.Errorf("fail: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cartItemWithProduct CartItemWithProduct
		if err := rows.Scan(
			&cartItemWithProduct.ProductId,
		); err != nil {
			return 0, fmt.Errorf("fail: %w", err)
		}
		cartItems = append(cartItems, &cartItemWithProduct)
	}
	return len(cartItems), nil
}

func (cis *CartItemService) UpdateQuantity(updatedCartItem *CartItemWithProduct, userId uuid.UUID) (bool, error) {
	_, err := cis.DB.Exec(`
		UPDATE cart_items
		SET quantity = $1
		WHERE user_id = $2 AND product_id = $3
	`, updatedCartItem.Quantity, userId, updatedCartItem.ProductId)
	if err != nil {
		return false, fmt.Errorf("fail: %w", err)
	}

	return true, nil
}

func (cis *CartItemService) Delete(productId string, user_id string) (bool, error) {
	_, err := cis.DB.Exec(`
		DELETE FROM cart_items
		WHERE product_id = $1 AND user_id = $2
	`, productId, user_id)
	if err != nil {
		return false, fmt.Errorf("fail: %w", err)
	}

	return true, nil
}
