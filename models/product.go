package models

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type ProductService struct {
	DB *sql.DB
}

type Product struct {
	ProductId          uuid.UUID
	ProductName        string
	ProductDescription string
	ImageURL           string
	Stock              int
	Price              int
	Featured           bool
	CategoryId         uuid.UUID
}

func (ps ProductService) Add(newProduct *Product) (*Product, error) {
	_, err := ps.DB.Exec(`
		INSERT INTO products (
			product_id, 
			product_name, 
			product_description, 
			image_url, 
			stock, 
			price, 
			featured, 
			category_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, newProduct.ProductId,
		newProduct.ProductName,
		newProduct.ProductDescription,
		newProduct.ImageURL,
		newProduct.Stock,
		newProduct.Price,
		newProduct.Featured,
		newProduct.CategoryId,
	)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	return newProduct, nil
}
