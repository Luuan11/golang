package main

import "fmt"

//Funçao init tem prioridade sobre a main, indepedente da ordem
func init() {
	fmt.Println("Prioridade na funcao init!")
}

func main() {
	fmt.Println("Prioridade na funcao main!")
}