package main

import "fmt"

func generico(inter interface{}){
	fmt.Println(inter)
}

// aceita todos os tipos
func main() {
	generico("string")
	generico(1)
	generico(true)
}