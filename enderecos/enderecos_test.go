package enderecos

import "testing"

func TestTipoEnderecos(t *testing.T) {
	// chamar funcao a ser testada e avaliar seu retorno
	input := "rua Paulista"
	sutRecebido := TipoEnderecos(input)
	expected := "Rua"

	if sutRecebido != expected {
		t.Error("QUEBROU  :: O Recebido é diferente do esperado")
	}
}
