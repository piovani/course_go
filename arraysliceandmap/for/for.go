package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, numeros := range numeros {
		fmt.Printf("%d) %d\n", i+1, numeros)
	}
}