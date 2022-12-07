package main

import "fmt"

func main() {
	
	// declaracao_com_inferencia com o operador "foquinha"  :=
	var1 := "Valor vars1"
	var2, var3 := true, 10	// declaracao mais_que_uma
	var4, var5, var6, var7 := "vTexto4", 555, false, "777" // declaracoes_multiplas

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)
	fmt.Println(var7)
}

/*
variaveis:
conceito_variavel: variavel é um recipiente onde se guarda uma informacao_valor.
como_nasce: ela nasce com valor inicial "nil / nulo" e apos alguma declaracao ela é registrada.

registro_nascimento_declaracao:
	 Uma variavel para existir tem que ser apresentada com uma declaracao que fala seu nome, seu tipo e valor , se nao ter um valor ela recebe por padrao um valor inicial da linguagem apra o tipo que ela percente.

	 tipo_do_valor: apos a primeira declaracao do tipo do valor nao se pode mudar mais o tipo do valor dela... posso mudar o conteudo mas o tipo nao mais.

	 formas_de_declaracoes:
	   declaracao_com_inferencia: nomeVar := valorInformacao

		 declaracao_verbosa:
		 informa com a clausula var o nome tipo e valor

		 declaracoes_multiplas:
		 declarando_por_inferencia mais_que_uma var na mesma linha ... na atribuicao ja passo os valores por esatr sendo por inferencia ..obs: estes valores tem que ser por justa posicao conforme as posicoes das variaveis da esquerda tem que combinar as posicoes...posso declarar mais que uma entao posso declarar e atribuir milhares infinitas.

	*/