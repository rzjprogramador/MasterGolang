package main

import "fmt"

func main() {
	usuarioMap := map[string]string{
		"nome":      "Reinaldo",
		"sobrenome": "Zacharias",
	}

	fmt.Println(usuarioMap)

	// acessar valor de chave : obs: somente apssando a chave como string
	fmt.Println(usuarioMap["nome"])
}

/*
CONCEITOS:
map é um objeto tipo json só que rigido.
declaracao: nomedoMap := map[tipoKey]tipoValue { props}
obs: todas as propriedades tem que ter virgula no final ... inclusive a ultima.

 - as chaves tem que ser todas do mesmo tipo declarado.
 - e os valores tem que ser todos do mesmo tipo declarado.
 exemplo se passou que as chaves serao string todas as chaves terao que ser strings
  se passou que os valores serao inteiro todas os valores terao que ser inteiros

	acessos:
	acessar_valor_de_chave : somente apssando a chave como string
	fmt.Println(usuarioMap["nome"])

*/
