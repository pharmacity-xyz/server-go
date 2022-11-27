package models

import (
	"database/sql"
	"fmt"
	"strings"

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
	hashedPassword, err := hash(password)
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}
	user.PasswordHash = hashedPassword
	_, err = us.DB.Exec(`
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
	if err != nil {
		return nil, fmt.Errorf("fail: %w", err)
	}
	return user, nil
}

func (us *UserService) Login(email, password string) (*User, error) {
	email = strings.ToLower(email)
	user := User{
		Email: email,
	}
	row := us.DB.QueryRow(`
		SELECT user_id, password_hash
		FROM users
		WHERE email=$1
	`, email)
	err := row.Scan(&user.UserId, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}
	return &user, nil
}

func (us *UserService) ChangePassword(userId uuid.UUID, newPassword string) (bool, error) {
	hashedPassword, err := hash(newPassword)
	if err != nil {
		return false, fmt.Errorf("fail: %w", err)
	}
	_, err = us.DB.Exec(`
		UPDATE users 
		SET password_hash = $1
		WHERE user_id = $2 
	`, hashedPassword, userId)
	if err != nil {
		return false, fmt.Errorf("fail: %w", err)
	}
	return true, nil
}

func hash(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("fail: %w", err)
	}
	return string(hashedBytes), nil
}
