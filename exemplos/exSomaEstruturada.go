package exemplo

import "fmt"

// type SumOut uint | error
// type SumOut uint error
// TODO NAO CONSEGUI RETORNAR MAIS QUE UM TIPO  >> NO TIPO PERSONALIZADO...

func Sum(a, b uint) uint {
	result, err := useSum(a, b)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func useSum(a, b uint) (uint, error) {
	if a+b > 100 {
		return 0, fmt.Errorf("Resultado nao pode ser maior que 100")
	}
	return a + b, nil
}

/*
implementacao:

resumo: fazer 2 acoes: [
	1- a logicaPrivada para moldar falha e sucesso,
	2- a executadoraPublicaRespondedora  que vai dar a resposta de falha e sucesso
	 executando a logica() capturando o sucesso e erro ,
	 verificando se tem erro responde este erro e depois responde tbm sucesso
	]

- fazer acaoLocalLogica() ::
que sera a logica, tera a regra de negocio, avaliada senao passar na regra vai retornar um feedback de excessao definido aqui....
ou se der tudo certo na regra vai retornar o objetivo ou erro vazio

- fazer AcaoPrincipal que vai ser compartilhada... e retornar os capturados erro ou sucesso

retornando_erro_ou_sucesso_da_subAcao:
 nela vamos pré executar a subAcao :: acaoLocalLogica()
- capturamos os resultados desta acao  :: capturando seu resultadoSucesso e erro .. tudo em variaveis

retornando_os_capturados:
fastFail:: se o erro tiver preenchido retornamos o erro capturado em uma saida de console.
- se nao tiver erro retornammos o resultado capturado

*/
