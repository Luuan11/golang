package main

import "fmt"

//criando uma struct
type user struct {
	name string
	age  uint
	address address
}

type address struct {
	street string
	number int
}

func main() {
	fmt.Println("Arquivo structs")

	var u user
	u.name = "Luan"
	u.age = 21
	fmt.Println(u)

	addressExample := address{"Rua dos bobos", 0}

	usuario2 := user{"Luan2", 21, addressExample}
	fmt.Println(usuario2)

	usuario3 := user{name: "Luan3"} //or {age: 21}
	fmt.Println(usuario3)
}
