package main

import "fmt"

func main() {
	VariarEmString()
}

func VariarEmString() {
	texto := "TEXTO"
	for indice, letra := range texto {
		// VAI PRINTAR EM TABELA ASC
		// fmt.Println(indice, letra)

		// VAI PRINTAR EM STRING
		fmt.Println(indice, string(letra))
	}
}

/*

# Loops For com Range

significado: Variar Para ...

restricoes: NAO DÁ PARA FAZER FOR COM RANGE EM STRUCTS ESTRUTURAS.

-  para ocultar a var que mostra o indice coloque um anderline no lugar.

 */
