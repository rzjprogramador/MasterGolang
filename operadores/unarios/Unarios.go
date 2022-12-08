package main

import "fmt"

func main() {
	aumentar := 10
	aumentar++
	aumentar++
	aumentar++

	diminuir := 20
	diminuir--
	diminuir--

	fmt.Println(aumentar)
	fmt.Println(diminuir)
}

/*
OPERADORES UNARIOS
entrada:
	descricao: operadores curtos que agem em uma referencia por vez.
	aplicabilidade: Aumentar  ou Diminuir o valor de uma referencia numero "INCREMENTAR"
	existentes: sao eles ++ para aumentar , -- para diminuir

processamento:
 acao: os encurtadores unarios aumentam ou diminuem de 1 em 1 unidade.
 abreviamos para numero++  ,  SERIA O MESMO DE numero = numero + 1 ,
 narracao: "numero é igual o que ela ja tem + mais 1"

saida:
++ para aumentar , -- para diminuir

*/
