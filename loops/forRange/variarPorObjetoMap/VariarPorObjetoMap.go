package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Reinaldo",
		"sobrenome": "Zacharias",
	}
	// fmt.Println(usuario)

	/*
		// iterar sobre o objMap
		for chave, valor := range usuario {
			fmt.Println(chave, valor)
		}
	*/

	// mostrar somente os valores
	for _, valor := range usuario {
		fmt.Println(valor)
	}

	// OBS: NAO DÁ PARA FAZER RANGE EM STRUCTS
}
