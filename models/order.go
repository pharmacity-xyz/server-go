package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pharmacity-xyz/server-go/responses"
)

type OrderService struct {
	DB *sql.DB
}

type Order struct {
	OrderId     uuid.UUID
	UserId      uuid.UUID
	TotalPrice  float64
	ShipAddress string
	OrderDate   time.Time
	ShippedDate time.Time
}

func (o OrderService) GetOrders() ([]*responses.OrderOverviewResponse, error) {
	// response := responses.OrderOverviewResponse

	return nil, nil
}

func (os OrderService) PlaceOrder(products []*CartItemWithProduct, userId uuid.UUID) (bool, error) {
	var totalPrice float64

	for i := 0; i < len(products); i++ {
		totalPrice += products[i].Price * float64(products[i].Quantity)
	}

	order := Order{
		OrderId:     uuid.New(),
		UserId:      userId,
		TotalPrice:  totalPrice,
		ShipAddress: "Tokyo",
		OrderDate:   time.Now(),
		ShippedDate: time.Now(),
	}

	_, err := os.DB.Exec(`
		INSERT INTO orders (order_id, user_id, total_price, ship_address, order_date, shipped_date)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, order.OrderId,
		order.UserId,
		order.TotalPrice,
		order.ShipAddress,
		order.OrderDate,
		order.ShippedDate)
	if err != nil {
		return false, fmt.Errorf("fail: %w", err)
	}

	for i := 0; i < len(products); i++ {
		_, err := os.DB.Exec(`
		INSERT INTO order_items (order_id, product_id, quantity, total_price)
		   VALUES ($1, $2, $3, $4)
		`, order.OrderId,
			products[i].ProductId,
			products[i].Quantity,
			products[i].Price*float64(products[i].Quantity),
		)
		if err != nil {
			return false, fmt.Errorf("fail: %w", err)
		}
	}
	_, err = os.DB.Exec(`
		DELETE FROM cart_items
		WHERE user_id = $1
	`, userId)
	if err != nil {
		return false, fmt.Errorf("fail: %w", err)
	}

	return true, nil
}
