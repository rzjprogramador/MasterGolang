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
