package main

import "fmt"

type user struct {
	nome  string
	idade uint
}

func (u user) salvar() {
	fmt.Printf("Salvando esse user: %s no banco da idade %d", u.nome, u.idade)
}

func main() {
	user1 := user{"Luan", 21}
	fmt.Println(user1)
	user1.salvar()
}