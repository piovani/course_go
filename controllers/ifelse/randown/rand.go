package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numRand() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	fmt.Printf("Numero %d", numRand())
}
