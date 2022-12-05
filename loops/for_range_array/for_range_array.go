package main

import "fmt"

func main() {
	items := []string{"Reinaldo", "Renata", "Guga"}

	printarItems(items)

}

func printarItems(array []string) {
	for _, valor := range array {
		fmt.Println(valor)
	}
}

/*
* para ocultar a var que mostra o indice coloque um anderline no lugar.
 */

// func retornarItems(array [3]string) {
// 	for _, valor := range array {
// 		valor
// 	}
// }
