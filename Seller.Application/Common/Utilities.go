package Common

import (
	"math/rand"
	"time"
)

func GetRandomCode(min int, max int) int {
	// set seed
	rand.Seed(time.Now().UnixNano())
	// generate random number and print on console
	return rand.Intn(max-min) + min
}
