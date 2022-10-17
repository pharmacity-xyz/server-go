package helpers

import (
	"math/rand"
)

func RandomNumber(n int) int {
	rand.Seed(43)
	value := rand.Intn(n)
	return value
}
