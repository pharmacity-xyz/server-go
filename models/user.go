package models

import (
	"database/sql"

	"github.com/google/uuid"
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

func (us *UserService) Register(data struct{Email string FirstName string}) (*User, error) {
		
	return nil, nil
}
