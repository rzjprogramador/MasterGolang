package main

import "fmt"

func main() {
	items := []string{"Reinaldo", "Renata", "Guga"}

	PrintarItems(items)

}

func PrintarItems(array []string) {
	for _, valor := range array {
		fmt.Println(valor)
	}
}

/*

# Loops For com Range

significado: Variar Para ...

restricoes: NAO DÁ PARA FAZER FOR COM RANGE EM STRUCTS ESTRUTURAS.

-  para ocultar a var que mostra o indice coloque um anderline no lugar.

 */

