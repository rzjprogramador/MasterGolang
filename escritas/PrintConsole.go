package escritas

import "fmt"

func PrintConsole() {
	texto := "Meu 1º texto"

	print(texto)

	fmt.Println(texto)
	
	fmt.Printf("Meu texto é >> %s ...", texto)

}

/*
// funcao print , sozinha só printa nao pula linha no console.

// funcao Println do obj fmt ela printa e pula linha no console

// Printf() do obj fmt ... posso renderizar variaveis atraves da tag %s porcentagemTipo.

*/