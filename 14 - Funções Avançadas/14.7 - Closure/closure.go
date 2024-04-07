package main

import "fmt"

func closure() func() {
	txt := "Texto rodando na funcao closure"

	funcaoClosure := func() {
		fmt.Println(txt)
	}
	return funcaoClosure
}

func main() {
	txt := "Texto rodando na funcao main"
	fmt.Println(txt)

	funcaoNova := closure()
	funcaoNova()
}