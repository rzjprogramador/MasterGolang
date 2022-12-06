package main

import (
	"fmt"
	"strconv"
)

func main() {
	texto := "10"
	resultConvertido, _ := strconv.ParseUint(texto, 10, 64)
	fmt.Println(resultConvertido)
}

/*
	* CONVERSAO STRING PARA NUMERO
	tutorial: https://pkg.go.dev/strconv
*/
