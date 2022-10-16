package main

import "log"

var s = "seven"

func main() {
	var s2 = "six"
	s := "eight"

	log.Println("s is ", s)
	log.Println("s2 is ", s2)

	saySomething(s)
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
