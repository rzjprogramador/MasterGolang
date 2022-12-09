package agrupamento

import (
	"fmt"
	"reflect"
)

func SliceArrayLivre() {
	var slice1 []string
	fmt.Println(slice1)

	sliceString := []string{"um", "dois", "tres"}
	fmt.Println(sliceString)

	sliceNumeros := []int32{2, 3, 4}
	fmt.Println(sliceNumeros)

	// Ver tipos com a libGO reflect
	fmt.Println(reflect.TypeOf(sliceString)) // []string

	// slicePopular[0]["quatro"]
	// fmt.Println(slicePopular)
}
