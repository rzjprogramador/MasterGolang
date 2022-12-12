package ferramentasFuncoes

import "log"

func recuperaExecucao() {
	if r := recover(); r != nil {
		log.Println("Recuperando execucao")
	}
}

func hasAprovado(num1, num2 float64) bool {
	defer recuperaExecucao()

	media := (num1 + num2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MEDIA É EXATAMENTE 6")
}

func ExecutePanic() {
	log.Println(hasAprovado(6, 6))
	log.Println("Pos execucao")
}

/*
resumo: panic() vai estourar erro caso o resultado da funcao seja exatamente o que nao quero ... recover() vai avaliar se tem isto que nao quero como resultado e vai recuperar o programa retornando a resposta que quero para a recuperacao do fluxo do programa.

panic:
panic sera uma excessao se a funcao der o resultado especifico que nao quero ela entra em panico como se fosse um throw e interrompe a execucao e o programa é abortado ele para entra em panico. se eu definir que tera panico caso isto acontecer com panic()

ele em panico vai chamar qualquer funcao que tenha sidoa diada com defer

entao no defer chamo uma funcao para recuperar a execucao do programa.

nesta recupera a execucao terei um recover() restauracao onde :
- SE a variavel auxiliar "r" que será recover() for diferente de nulo é porque esta com valor em panico...entao vai recuperar a funcao em panico.

*/
