package main

import (
	"fmt"
)

func main() {

	var nome string
	var sobrenome string

	fmt.Println("Digite seu nome: ")
	fmt.Scan(&nome, &sobrenome)

	fmt.Println(nome)
	fmt.Println(sobrenome)
}
