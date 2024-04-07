package main

import "fmt"

type user struct {
	nome  string
	idade uint
}

func (u user) salvar() {
	fmt.Printf("Salvando esse user %s da idade %d no banco de dados de usuarios \n", u.nome, u.idade)
}

func (u user) verificaIdade(){
	if u.idade >= 18{
		fmt.Printf("informação extra, o user %s é maior de idade \n", u.nome)
	} else {
		fmt.Printf("que pena, o user %s é menor de idade \n", u.nome)
	}
}

func main() {
	user1 := user{"Luan", 21}
	
	user2 := user{
		nome: "Giovana",
		idade: 22,
	}

	user3 := user{"Alvaro", 17}
	
	user1.salvar()
	user2.salvar()
	user3.salvar()
	user1.verificaIdade()
	user2.verificaIdade()
	user3.verificaIdade()
}