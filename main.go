package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	myVar := "hello"

	switch myVar {
	case "cat":
		log.Println("cat")

	case "dog":
		log.Println("dog")

	case "fish":
		log.Println("fish")

	default:
		log.Println("does not match")
	}
}

func saySomething(s string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s, "World"
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
