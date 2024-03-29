package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	age       int
	altura    uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa1 := pessoa{"Joao", "Carlos", 20, 178}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Sistemas de Informação", "Uninove"}
	fmt.Println(estudante1)
}
