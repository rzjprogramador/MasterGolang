package erros

import (
	"errors"
	"log"
)

func erroService(s string) (string, error) {
	valor := s
	if valor == "foo" {
		return "", errors.New("nao pode ser foo")
	}
	return valor, nil
}

func PrintErrosResolver() string {
	result, err := erroService("foo")

	if err != nil {
		// log.Fatalf("o erro foi :: %v", err)
		log.Fatal(err.Error())
	}
	return result
}

/*
Printar_erros:
	#preferencia::
		log.Fatal(err.Error())
		ou se preferir mais uma mensagem antes de printar o erro use:
		log.Fatalf ("o erro foi :: %v", variavelErro )
*/
