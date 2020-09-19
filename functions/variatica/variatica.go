package main

import "fmt"

func media(numeros ...float64) float64 {
	if len(numeros) == 0 {
		return 0
	}

	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	// fmt.Printf("MÉDIA: %.2f", media(7.7, 8.1, 5.9))
	fmt.Printf("MÉDIA: %.2f", media())
}
