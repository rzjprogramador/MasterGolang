package funcoes

import "fmt"

func MultiplasEntradas() {
	fmt.Println(reduceSomaFloat64(100, 100.1, 100.77))
}

func reduceSomaFloat64(numeros ...float64) float64 {
	total := 0.0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

/*
funcoes_variaticas:
  conceito: pode receber diversos params, sendo assim sempre terá um slice que podera ser iterado e resultar um valor reduzido deste slice no caso da entrada ser numerica.
	podemos fazer um semelhante reduce do js.
	obs: a variavel de retorno modificada tera que ter o tipo de retorno valido da funcao.
	ex: se a var de retorno for um float64 a entrada tera que ser tambem float64

	declaracao: para ter multiplas entradas acrescentamos ... "3 pontos ao tipo da entrada"


*/
