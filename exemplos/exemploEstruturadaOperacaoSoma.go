package exemplo

import "fmt"

func ExemploEstruturadaOperacaoSoma(a, b int) int {
	result, err := soma(a, b)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func soma(a, b int) (int, error) {
	if a+b > 100 {
		return 0, fmt.Errorf("Resultado nao pode ser maior que 100")
	}
	return a + b, nil
}

/*
implementacao:

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
