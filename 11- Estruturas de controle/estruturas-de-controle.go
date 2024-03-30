package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	num := 10

	if num > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//criando a variavel dentro de um if
	//variavel limitada a uso somente ao escopo do if
	if outroNum := num; outroNum > 0{
		fmt.Println("Numero é maior que zero")
	} else if num < - 10 {
		fmt.Println("Número é menor que -10")
	} else{
		fmt.Println("Entre 0 e -10")
	}

	
}
