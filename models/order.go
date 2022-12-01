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

func (os OrderService) GetOrders(userId uuid.UUID) (*[]responses.OrderOverviewResponse, error) {
	response := []responses.OrderOverviewResponse{}

	rows, err := os.DB.Query(`
		SELECT orders.order_id, order_date, orders.total_price, product_name, image_url FROM orders
		JOIN order_items ON order_items.order_id = orders.order_id
		JOIN products ON products.product_id = order_items.product_id
		WHERE user_id = $1
	`, userId)

	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	defer rows.Close()

	orderId := uuid.UUID{}

	for rows.Next() {
		var orderOverviewRes responses.OrderOverviewResponse
		if err := rows.Scan(
			&orderOverviewRes.OrderId,
			&orderOverviewRes.OrderDate,
			&orderOverviewRes.TotalPrice,
			&orderOverviewRes.Product,
			&orderOverviewRes.ProductImageUrl,
		); err != nil {
			return nil, fmt.Errorf("fail: %w", err)
		}
		if orderId == orderOverviewRes.OrderId {
			continue
		}
		orderId = orderOverviewRes.OrderId
		response = append(response, orderOverviewRes)
	}

	return &response, nil
}

type OrderDetail struct {
	OrderDate           time.Time
	OrderTotalPrice     float64
	ProductId           uuid.UUID
	ProductName         string
	ImageUrl            string
	Quantity            int64
	OrderItemTotalPrice float64
}

func (os OrderService) GetOrderDetails(userId uuid.UUID, orderId string) (*responses.OrderDetailsResponse, error) {
	var response responses.OrderDetailsResponse

	rows, err := os.DB.Query(`
		SELECT order_date, orders.total_price, order_items.product_id, product_name, image_url, order_items.quantity, order_items.total_price 
		FROM orders
		inner JOIN order_items ON order_items.order_id = orders.order_id
		JOIN products ON products.product_id = order_items.product_id
		WHERE user_id = $1 AND orders.order_id = $2
	`, userId, orderId)

	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	defer rows.Close()

	var orderDetail OrderDetail

	for rows.Next() {
		if err := rows.Scan(
			&orderDetail.OrderDate,
			&orderDetail.OrderTotalPrice,
			&orderDetail.ProductId,
			&orderDetail.ProductName,
			&orderDetail.ImageUrl,
			&orderDetail.Quantity,
			&orderDetail.OrderItemTotalPrice,
		); err != nil {
			return nil, fmt.Errorf("fail: %w", err)
		}
		orderDetailProductResponse := responses.OrderDetailsProductResponse{
			ProductId:   orderDetail.ProductId,
			ProductName: orderDetail.ProductName,
			ImageUrl:    orderDetail.ImageUrl,
			Quantity:    orderDetail.Quantity,
			TotalPrice:  orderDetail.OrderItemTotalPrice,
		}
		response.OrderDetailsProductResponse = append(response.OrderDetailsProductResponse, orderDetailProductResponse)
	}

	response.OrderDate = orderDetail.OrderDate
	response.TotalPrice = orderDetail.OrderTotalPrice

	return &response, nil
}

func (os OrderService) GetOrdersForAdmin() (*[]responses.OrderOverviewResponse, error) {
	response := []responses.OrderOverviewResponse{}

	rows, err := os.DB.Query(`
		SELECT orders.order_id, order_date, orders.total_price, product_name, image_url FROM orders
		JOIN order_items ON order_items.order_id = orders.order_id
		JOIN products ON products.product_id = order_items.product_id
	`)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}
	defer rows.Close()

	orderId := uuid.UUID{}

	for rows.Next() {
		var orderOverviewRes responses.OrderOverviewResponse
		if err := rows.Scan(
			&orderOverviewRes.OrderId,
			&orderOverviewRes.OrderDate,
			&orderOverviewRes.TotalPrice,
			&orderOverviewRes.Product,
			&orderOverviewRes.ProductImageUrl,
		); err != nil {
			return nil, fmt.Errorf("fail: %w", err)
		}
		if orderId == orderOverviewRes.OrderId {
			continue
		}
		orderId = orderOverviewRes.OrderId
		response = append(response, orderOverviewRes)
	}

	return &response, nil
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
