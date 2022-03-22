package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta feira"
	case 7:
		return "Sabado"
	default:
		return "numero invalido"
	}

}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sabado"
	default:
		return "numero invalido"
	}
}

func main() {
	fmt.Println("Switch!")

	dia := diaDaSemana(3)
	fmt.Println(dia)

}
