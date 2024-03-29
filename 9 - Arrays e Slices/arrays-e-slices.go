package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	//Nesse exemplo aqui, usamos os ... para que o array defina o tamanho dele, sendo assim mais flexivel
	//Não é um tamanho dinamico, mas sim, flexivel ao valor que ele vai receber
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array3)

	//similar ao array, porem esse sim é dinamico
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	//para retornamos os tipos do slice e do array3
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//adiciona o item no slice selecionado e retorna um novo slice com o valor inserido
	slice = append(slice, 18)
	fmt.Println(slice)

	//pegando uma fatia do array, do index 1 ao 3 posição
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "posição alterada"
	fmt.Println(slice2)
}
