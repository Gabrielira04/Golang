package main

import "fmt"

func main() {
	fmt.Println("estruturas de controle")

	numero := -5

	if numero > 15 {
		fmt.Println("é maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("nuemro é maior que zero")
	} else if numero < -10 {
		fmt.Println("nuemro menor que -10")
	} else {
		fmt.Println("Está entre 0 e -10")
	}

}
