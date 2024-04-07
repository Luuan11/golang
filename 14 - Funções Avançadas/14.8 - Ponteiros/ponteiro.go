package main

import "fmt"

func inverterNum(num int) int {
	return num * -1
}

func inverterNumPonteiro(num *int){
	*num = *num * -1
}

func main() {
	num := 20
	numInvertido := inverterNum(num)

	//podemos notar que sem a utilização de ponteiros, o valor da variavel não foi alterado, ainda consta como 20.
	fmt.Println(numInvertido)
	fmt.Println(num)

	println("===== COM PONTEIRO =====")

	numInvertidoPonteiro := 40
	inverterNumPonteiro(&numInvertidoPonteiro)
	fmt.Println(numInvertidoPonteiro)
}