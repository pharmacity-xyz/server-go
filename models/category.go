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

func (cs *CategoryService) GetAll() ([]*Category, error) {
	var categories []*Category
	rows, err := cs.DB.Query(`
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

func (cs *CategoryService) Add(newCategory *Category) (*Category, error) {
	_, err := cs.DB.Exec(`
		INSERT INTO categories (category_id, name)
		VALUES ($1, $2)
	`, newCategory.CategoryId, newCategory.Name)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	return newCategory, nil
}

func (cs *CategoryService) Update(updatedCategory *Category) (*Category, error) {
	_, err := cs.DB.Exec(`
		UPDATE categories
		SET name = $1
		WHERE category_id = $2
	`, updatedCategory.Name, updatedCategory.CategoryId)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}

	return updatedCategory, nil
}

func (cs *CategoryService) Delete(categoryId string) error {
	_, err := cs.DB.Exec(`
		DELETE FROM categories
		WHERE category_id = $1
	`, categoryId)
	if err != nil {
		return fmt.Errorf("fail: %w", err)
	}

	return nil
}
