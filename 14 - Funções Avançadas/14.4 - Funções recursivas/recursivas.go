package main

import "fmt"

func main() {
	fmt.Println("Funções recursivas")
	posicao := uint(12)
	fmt.Println(fibonacci(posicao))
}

func fibonacci(posicao uint) uint{
	if posicao <= 1{
		return	posicao
	}

	return fibonacci(posicao -2) + fibonacci(posicao-1)
}