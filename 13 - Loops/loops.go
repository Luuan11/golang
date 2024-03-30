package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	//incrementar até o valor chegar a 10.
	for i < 10 {
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	//definindo a variavel já no for
	for j := 0; j < 10; j+=2 {
		fmt.Println("Incrementando J", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	//percorrendo todo o array e lançando posição e valor
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//navegando pela palavra "Palavra"
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	user := map[string]string{
		"nome":      "Luan",
		"Sobrenome": "Fernando",
	}

	//navegando pelo Map
	for chave, valor := range user {
		fmt.Println(chave, valor)

	}

	type userStruct struct {
		nome string
		sobrenome string
	}

	//Em Struct, já não é possível
	/*
		user2 := userStruct{"Luan", "Chaves"}
		for chave, valor := range user2 {
			fmt.Println(chave, valor)
		}
	*/
}
