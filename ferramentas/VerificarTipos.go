package ferramentas

import (
	"fmt"
	"reflect"
)

func VerificarTipos() {
	texto := "foo"
	numero := 10
	grupoString := []string{"foo", "bar"}

	// Verificar tipos com a libGO reflect usando dela o metodo TypeOf()

	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(grupoString))
}
