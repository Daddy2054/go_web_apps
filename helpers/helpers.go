package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}
