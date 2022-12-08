package main

import "fmt"

func main() {
	aumentar := 10
	aumentar++
	aumentar++
	aumentar += 55

	diminuir := 20
	diminuir--
	diminuir--
	diminuir -= 10

	fmt.Println(aumentar) // comecou com 10, aumentou 2 vezes de 1 deu 12, aumentou 55, total: 67
	fmt.Println(diminuir) // comecou com 20, diminuiu 2 vezes de 1 deu 18, diminuiu -10 total: 8
}

/*
OPERADORES UNARIOS
entrada:
	descricao: operadores curtos que agem em uma referencia por vez.
	aplicabilidade: Aumentar  ou Diminuir o valor de uma referencia numero "INCREMENTAR"
	existentes: sao eles ++ para aumentar , -- para diminuir

processamento:
 acao: os encurtadores unarios aumentam ou diminuem :

 narracao: "numero é igual o que ela ja tem + mais 1 ..ou -1 se for dinuir"

 ++ aumenta de 1 em 1 unidade. // SERIA O MESMO DE numero = numero + 1 ,
 -- diminui de 1 em 1 unidade // SERIA O MESMO DE numero = numero -1,
 += 15  :: vai aumentar o que ja tem para mais 15
 -= 33  :: vai diminuir o que tem por 33


saida:
++ para aumentar , -- para diminuir

*/
