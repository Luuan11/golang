package main

import "fmt"

func recuperacao(){
	if r := recover(); r != nil{
		fmt.Println("Recuperado")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperacao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Media é 6!")
}

func main() {
	fmt.Println(alunoAprovado(8,6))
	fmt.Println("Fim de exec")
}