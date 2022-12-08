package escritas

import "fmt"

func PrintConsole() {
	texto := "Meu 1º texto"

	// funcao print , sozinha só printa nao pula linha no console.
	print(texto)

	// funcao Println do obj fmt ela printa e pula linha no console
	fmt.Println(texto)

	// Printf() do obj fmt ... posso renderizar variaveis atraves da tag %s porcentagemTipo.
	fmt.Printf("Meu texto é >> %s ...", texto)

}
