package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}

	fmt.Printf("A1: %d, S1: %d\n", a1, s1)

	a2 := a1[1:3]

	fmt.Printf("A2: %d", a2)
}
