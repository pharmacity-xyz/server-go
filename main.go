package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       uint
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@gmail.com", 14})
	users = append(users, User{"John", "Smith", "john@gmail.com", 14})
	users = append(users, User{"John", "Smith", "john@gmail.com", 14})
	users = append(users, User{"John", "Smith", "john@gmail.com", 14})

	for _, l := range users {
		log.Println(l.FirstName)
	}
}
