package main

import "fmt"

func funcao1() {
	fmt.Println("Escrevendo")
}

func funcao2(){
	fmt.Println("Escrevendo novamente")
}

func alunoAprovado(n1, n2 float64) bool{
	defer fmt.Println("Media calculada, resultado está sendo processado!")
	fmt.Println("Media calculada, resultado está sendo processado!")
	media := (n1 + n2) / 2

	if media >= 6{
		return true
	}
	return false
}

func main() {
	defer funcao1() // adiando a função
	funcao2()

	alunoAprovado(10,10)
}