package helpers

import "math/rand"

func RandomRange(min int, max int) int {
	return min + rand.Intn(max-min)
}
