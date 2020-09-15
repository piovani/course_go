package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[17938535095] = "MARIA"
	aprovados[48454718095] = "JOSE"
	aprovados[72153031010] = "FERNANDA"

	fmt.Println(aprovados)

	for cpf, name := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}
}
