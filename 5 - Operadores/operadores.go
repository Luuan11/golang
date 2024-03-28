package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// não é possivel somar valores que sejam de int diferentes
	var num1 int16 = 10
	var num2 int16 = 25
	soma2 := num1 + num2
	fmt.Println(soma2)

	//atribuicao
	var variavel1 string = "String"
	variavel2 := "String 2"

	fmt.Println(variavel1, variavel2)

	//Operadores
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//Operadores Logicos
	fmt.Println(false && true)  //se todos forem verdade, passa
	fmt.Println(true || false) //se um for verdade, passa
	fmt.Println(!true)         //para negar o valor, se for true, negando vira false

	//Operadores unarios
	numeroTeste := 10
	numeroTeste++ 
	numeroTeste += 15 // numeroTeste = numertoTeste + 15
	numeroTeste-- 
	numeroTeste -=20  // numeroTeste = numertoTeste - 20
	numeroTeste *=3 //  numeroTeste = numertoTeste * 3
	numeroTeste /= 10
	numeroTeste %= 3

	fmt.Println(numeroTeste)

	//Operador Tenario || NAO EXISTE EM GOLANG
	// textoStr := num > 5 ? "Maior que 5" : "Menor que 5"
	var texto string
	if numeroTeste > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
