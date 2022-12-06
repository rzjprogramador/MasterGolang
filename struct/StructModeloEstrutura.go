package main

import "fmt"

/*
* conceito: struct é tipo uma classe com new : é um modelo em tipagem de um objeto entidade.
* obs: em Go nao tem classes esta estrutura é o que mais se assemelha.
*/

// declaracao struct : a declaracao pode ser feita fora do escopo da funcao.

type Usuario struct {
	nome      string
	sobrenome string
	idade     uint8
	estado    bool
}

func main() {

	// struct formaDeUso1
	usuario1 := Usuario{
		nome:      "Reinaldo",
		sobrenome: "Zacharias",
		idade:     45,
		estado:    true,
	}

	// struct formaDeUso2
	var usuario2 Usuario
	usuario2.nome = "Renata"
	usuario2.idade = 40
	/*
		obs: senao passar todos valores da estrutura ele devolve os valores inicias default das props que estao faltando.
	*/

	fmt.Println(usuario1) // outup: {Reinaldo Zacharias 45 true}

	fmt.Println(usuario2) // output: {Renata "" 40 false}
}
