package main

import (
	"log"

	"github.com/pharmacity-xyz/server/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Hello"
	log.Println(myVar.TypeName)
}
