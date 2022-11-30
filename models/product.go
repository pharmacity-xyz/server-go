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

func (ps ProductService) GetAll() ([]*Product, error) {
	var products []*Product
	rows, err := ps.DB.Query(`
		SELECT *
		FROM products
	`)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.ProductDescription,
			&product.ImageURL,
			&product.Stock,
			&product.Price,
			&product.Featured,
			&product.CategoryId,
		); err != nil {
			return nil, fmt.Errorf("fail: %w", err)
		}
		products = append(products, &product)
	}
	return products, nil
}

func (ps ProductService) GetProductByProductId(productId string) (*Product, error) {
	var product Product

	row := ps.DB.QueryRow(`
		SELECT *
		FROM products
		WHERE product_id = $1
	`, productId)
	err := row.Scan(&product.ProductId,
		&product.ProductName,
		&product.ProductDescription,
		&product.ImageURL,
		&product.Stock,
		&product.Price,
		&product.Featured,
		&product.CategoryId)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	return &product, nil
}

func (ps ProductService) GetProductByCategoryId(categoryId string) ([]*Product, error) {
	var products []*Product

	rows, err := ps.DB.Query(`
		SELECT * 
		FROM products
		WHERE category_id = $1
	`, categoryId)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	for rows.Next() {
		var product Product
		if err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.ProductDescription,
			&product.ImageURL,
			&product.Stock,
			&product.Price,
			&product.Featured,
			&product.CategoryId,
		); err != nil {
			return nil, fmt.Errorf("fail: %w", err)
		}
		products = append(products, &product)
	}

	return products, nil
}

func (ps ProductService) Search(searchWord string) ([]*Product, error) {
	var products []*Product

	rows, err := ps.DB.Query(`
		SELECT * 
		FROM products
		WHERE LOWER(product_name) LIKE $1
	`, "%"+searchWord+"%")
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	for rows.Next() {
		var product Product
		if err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.ProductDescription,
			&product.ImageURL,
			&product.Stock,
			&product.Price,
			&product.Featured,
			&product.CategoryId,
		); err != nil {
			return nil, fmt.Errorf("fail: %w", err)
		}
		products = append(products, &product)
	}

	return products, nil
}

func (ps ProductService) FeaturedProducts() ([]*Product, error) {
	var products []*Product

	rows, err := ps.DB.Query(`
		SELECT * 
		FROM products
		WHERE featured = true
	`)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	for rows.Next() {
		var product Product
		if err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.ProductDescription,
			&product.ImageURL,
			&product.Stock,
			&product.Price,
			&product.Featured,
			&product.CategoryId,
		); err != nil {
			return nil, fmt.Errorf("fail: %w", err)
		}
		products = append(products, &product)
	}

	return products, nil
}
