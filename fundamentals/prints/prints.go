package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print("Linha.")

	x := 3.141516

	// fmt.Println("Teste: " + x)

	xs := fmt.Sprint(x)
	fmt.Println("Teste: " + xs)
}
