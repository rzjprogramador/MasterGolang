package main

import (
	"fmt"
	"_fundamentos/modules/soma"
)

func main() {
	somado := soma.Somar(10, 10)

	fmt.Println("Estou no main...")
	fmt.Println("A soma é  :: ", somado)
}