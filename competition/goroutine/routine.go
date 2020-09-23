package main

import (
	"fmt"
	"time"
)

func fale(pessoa, text string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, text, i+1)
	}
}

func main() {
	// fale("MARIA", "Ola", 3)
	// fale("JOAO", "So posso falar depois de voce", 1)

	// go fale("Maria", "Ei...", 500)
	// go fale("JOAO", "Opa...", 500)

	go fale("MARIA", "Endendi", 10)
	fale("JOAO", "PARABENS", 5 )

	fmt.Println("FIM")
}
