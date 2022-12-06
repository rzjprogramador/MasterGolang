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
