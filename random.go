package utils

import "math/rand"

// -> crypto/rand
func RandInt(min int, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}
