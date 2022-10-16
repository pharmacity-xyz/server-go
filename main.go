package main

import (
	"log"
	"sort"
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
	var mySlice []int

	mySlice = append(mySlice, 5)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)

	log.Println(mySlice)
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
