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
	OrderItems  []OrderItem
}

func (o OrderService) GetOrders() ([]*responses.OrderOverviewResponse, error) {
	response := responses.OrderOverviewResponse
}

func (os OrderService) PlaceOrder(products []*CartItemWithProduct, userId uuid.UUID) (bool, error) {
	var totalPrice float64

	for i := 0; i < len(products); i++ {
		totalPrice += products[i].Price * float64(products[i].Quantity)
	}

	var orderItems []OrderItem
	for i := 0; i < len(products); i++ {
		orderItems = append(orderItems, OrderItem{
			ProductId:  products[i].ProductId,
			Quantity:   products[i].Quantity,
			TotalPrice: totalPrice,
		})
	}

	order := Order{
		OrderId:     uuid.New(),
		UserId:      userId,
		TotalPrice:  totalPrice,
		ShipAddress: "Tokyo",
		OrderDate:   time.Now(),
		ShippedDate: time.Now(),
		OrderItems:  orderItems,
	}

	_, err := os.DB.Exec(`
		INSERT INTO orders (order_id, user_id, total_price, ship_address, order_data, shipped_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, order.OrderId, order.UserId, order.TotalPrice, order.ShipAddress, order.OrderDate, order.ShippedDate, order.OrderItems)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	return newCategory, nil

	return true, nil
}
