package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 10
	fmt.Println("i is set to", i)

	whatWasSaied := saySomething()

	fmt.Println("The function returned", whatWasSaied)
}

func saySomething() string {
	return "something"
}
