package main

import "fmt"

func main() {

	verdadeiro, falso := true, false 

	// OPERADORES LOGICOS ( && , || ) : SÓ RETORNAM BOLEANOS
	
	fmt.Println(verdadeiro && falso) // false

	// operador_E : && todas tem que ser verdade pra retornar verdade , narracao: Esta "E" estas pra ser verdade.

	fmt.Println(verdadeiro || falso) // true

	// operador_OU : || basta uma ser verdade pra retornar verdade, narracao: Pode ser Esta "OU" estas pra ser verdade.

}
