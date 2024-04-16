package generatives

import "math/rand"

func RandomNumbers() int {
	min := 12000
	max := 34000
	randNum := rand.Intn(max-min) + min
	return randNum
}
