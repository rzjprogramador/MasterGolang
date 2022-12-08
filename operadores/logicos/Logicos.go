package main

import "fmt"

func main() {

	verdadeiro, falso := true, false
	// foo := "BAR"

	/*
	 OPERADORES LOGICOS ( && , ||, ! ) :
	 processamento: SÃO PERGUNTAS , É VERDADEIRO (true) ? É FALSO (false)?

	 saida:
	 ASSIM QUE ENCONTRAM A PRIMEIRA OCORRENCIA DO SEU OBJETIVO JÁ RETORNAM BOLEANOS TRUE OU FALSE -- E INTERROMPEM AS VERIFICACOES
	*/

	fmt.Println(verdadeiro && falso) // false

	// operador_E : && todas tem que ser verdade pra retornar verdade , narracao: Esta "E" estas pra ser verdade.

	fmt.Println(verdadeiro || falso) // true

	// operador_OU : || basta uma ser verdade pra retornar verdade, narracao: Pode ser Esta "OU" estas pra ser verdade.

	fmt.Println(!verdadeiro)
	// operador_NEGACAO : ! inverte o valor bolaeano de uma referencia, se era true vira false e vice e versa

	// fmt.Println(!!foo) // TODO: TENTANDO INVERTER PARA BOLEANO UMA REFERENCIA NAO BOLEANA EM GO

}
