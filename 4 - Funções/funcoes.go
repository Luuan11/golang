package main

import "fmt"

func somar(num1 int, num2 int ) int{
	return num1 + num2
}

func calculosMath(n1, n2 int) (int, int){
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10 , 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	res := 	f("ola mundo, fui passado por parametros!")
	resSomar, resSubtracao := calculosMath(10,15)

	fmt.Println(res)
	fmt.Println(resSomar, resSubtracao)

	
}