package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	fmt.Println("heranca")

	p1 := pessoa{"joão", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "engenharia", "USP"}
	fmt.Println(e1)
}
