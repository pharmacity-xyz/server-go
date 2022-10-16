package main

import "log"

var s = "seven"

func main() {
	var s2 = "six"

	log.Println("s is ", s)
	log.Println("s2 is ", s2)
}

func saySomething(s string) (string, string) {
	return s, "World"
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
