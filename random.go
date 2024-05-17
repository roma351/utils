package utils

import (
	rand2 "crypto/rand"
	"log"
	"math/big"
	"math/rand"
)

func RandInt(min int, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

func RandInt2(min int, max int) int {
	if min == max {
		return min
	}
	n := max - min
	randInt, err := rand2.Int(rand2.Reader, big.NewInt(int64(n)))
	if err != nil {
		log.Fatalln("Error generating random number:", err)
	}

	return int(randInt.Int64()) + min
}
