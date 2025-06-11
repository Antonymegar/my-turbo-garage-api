package utils

import "math/rand"

// CreateRandomNumber ...
func CreateRandomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}
