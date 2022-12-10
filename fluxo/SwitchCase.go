package fluxo

import "fmt"

func SwitchCase() {
	fmt.Println(diaDaSemana((10)))
}

func diaDaSemana(numero uint8) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	}
	return "Default : entrada Invalida"
}

/*
- Obs: em GO nao temos o break para parar as avaliacoes

 Mundando o retorno conforme a entrada do utilizador:
 obs: assim como estou retornando uma string posso retornar qualquer coisa objeto, funcao etc.
 nesta implementacao a entrada do Utilizador sera comparada com o case.
 narracao: caso o numero seja 1 reotrna isto, caso seja 2 que entrar retorne isto...

*/
