package testes

import (
	"fmt"

	"github.com/MasterGolang/testes/endereco/paraTest"
)

func TestEndereco() {
	fnTipoEndereco := paraTest.TipoEnderecos("rua Reinaldo")
	fmt.Println(fnTipoEndereco)
}
