package main

import "fmt"

func main() {
	fmt.Println("Maps")
	//map[tipo das chaves]valores
	//precisam ser ambos do mesmo tipo
	user := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "carlos",
	}

	fmt.Println(user["nome"])

	user2 := map[string]map[string]string{
		"nome": {
			"Primeiro": "João",
			"Ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campos": "Campus 1",
		},
	}

	fmt.Println(user2)
	//Para deletar, fazemos dessa forma
	delete(user2, "nome")
	fmt.Println(user2)

	user2["signo"] = map[string]string{
		"nome": "Virgem",
	}

	fmt.Println(user2)

}
