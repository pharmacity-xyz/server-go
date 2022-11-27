package models

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId       uuid.UUID
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	City         string
	Country      string
	CompanyName  string
	Role         string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Register(user *User, password string) (*User, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}
	user.PasswordHash = string(hashedBytes)

	row := us.DB.QueryRow(`
		INSERT INTO users (user_id, email, password_hash, first_name, last_name, city, country, company_name, role)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
	`, user.UserId,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.City,
		user.Country,
		user.CompanyName,
		user.Role,
	)
	err = row.Scan()
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}
	return user, nil
}
