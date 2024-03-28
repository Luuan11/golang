package main

import "fmt"

func main() {
	num := -100000000
	fmt.Println(num)

	var num2 uint32 = 10000
	fmt.Println(num2)

	// alias
	//int32 = rune
	var num3 rune = 123456
	fmt.Println(num3)

	//byte = uint8
	var num4 byte = 123
	fmt.Println(num4)

	//números reais
	var numReal1 = 123.45
	fmt.Println(numReal1)
	
	var numReal2 = 12345.67
	fmt.Println(numReal2)

	//Strings
	var str string = "Exemplo"
	fmt.Println(str)

	str2 := "Exemplo2"
	fmt.Println(str2)

	// não há char em Golang

	//valor 0 em variavels
	var texto string
	fmt.Println(texto)
	//gera apenas um espaco em branco no console

	//booleanos
	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)
	//gera o valor nil
}