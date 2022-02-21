package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
	dadosAd  dadosAd
}

type endereco struct {
	logadouro string
	numero    uint8
}
type dadosAd struct {
	telefone     uint64
	escolaridade string
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos,", 0}

	dadoscomp := dadosAd{88442255, "Médio completo"}

	usuario2 := usuario{"Davi", 21, enderecoExemplo, dadoscomp}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)

}
