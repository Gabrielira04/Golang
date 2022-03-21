package main

import "fmt"

func main() {
	usuario := map[string]string{ //[tipo das chaves]Tipo dos valores
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":    "João",
			"ultimo nome": "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"Campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "escorpião",
	}
	fmt.Println(usuario2)
}
