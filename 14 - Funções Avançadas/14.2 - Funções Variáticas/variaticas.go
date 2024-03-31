package main

import "fmt"

// passando de 1 a n ints
func soma(num ...int) int {
	total := 0
	for _, numero := range num {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println(totalSoma)

	escrever("Ola mundo", 10, 2, 3, 4, 5)
}
