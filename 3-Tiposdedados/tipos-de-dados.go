package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -1000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias = apelido
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//NUMERO REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1234568785.45
	fmt.Println(numeroReal3)

	//FIM NUMEROS REAIS

	//STRING

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	//FIM STRING

	var booleano bool = true
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
