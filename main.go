package main

import (
	"fmt"
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
	myMap := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMap["me"] = me

	fmt.Println(myMap["me"].FirstName)

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
