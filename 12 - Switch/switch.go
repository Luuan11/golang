package main

import "fmt"

func diaDaSemana(num int) string {
	switch num {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "sabado"
	//caso o numero informado não esteja listado, passe a condição ao Default
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(num int) string {
	var diaDaSemana string

	switch {
	case num == 1:
		diaDaSemana = "domingo"
		//fallthrough pega o resultado esperado e passa para a proxima condição
		fallthrough
	case num == 2:
		diaDaSemana = "segunda"
	case num == 3:
		diaDaSemana = "terça"
	case num == 4:
		diaDaSemana = "quarta"
	case num == 5:
		diaDaSemana = "quinta"
	case num == 6:
		diaDaSemana = "sexta"
	case num == 7:
		diaDaSemana = "sabado"
	default:
		return "Número inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(200)
	fmt.Println(dia)
	fmt.Println("----------")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
