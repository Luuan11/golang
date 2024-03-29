package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var teste1 int = 10
	var teste2 int = teste1

	teste1++
	//apenas estará somando o valor do teste1, indo pra 11
	fmt.Println(teste1, teste2)

	//PONTEIRO, É UMA REFERENCIA DE MEMORIA
	var teste3 int 
	var ponteiro *int
	
	teste3 = 100
	//para pegar o endereço da memoria, se colocamos um *antes de defini la no println, pegaremos o valor
	ponteiro = &teste3

	fmt.Println(teste3, ponteiro)
}
