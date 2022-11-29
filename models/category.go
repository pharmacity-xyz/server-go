package models

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type CategoryService struct {
	DB *sql.DB
}

type Category struct {
	CategoryId uuid.UUID
	Name       string
}

func (us *CategoryService) GetAll() ([]*Category, error) {
	var categories []*Category
	rows, err := us.DB.Query(`
		SELECT *
		FROM categories
	`)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		if err := rows.Scan(
			&category.CategoryId,
			&category.Name,
		); err != nil {
			return nil, fmt.Errorf("fail: %w", err)
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
