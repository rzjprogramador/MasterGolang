package main

import "fmt"

func main() {
	repeteTexto()
}

func repeteTexto() {

	for varRepete := 0; varRepete <= 10; varRepete++ {
		fmt.Println("vai repetir este valor  numero de vezes passada na condicao: ", varRepete)
		fmt.Println(varRepete) // ver valor da varInteracao somente dentro do escopo
	}
}
