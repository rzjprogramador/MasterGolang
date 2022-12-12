package ferramentasFuncoes

import (
	"fmt"
	"log"
)

// --- exemplo 2 Defer Adiando

func AvaliacaoAlunoAdiando(n1, n2 int) bool {

	defer log.Println("Media calculada o resultado sera calculado ..esta msg é pra apos o retorno sendo true ou false")

	fmt.Println("Entrando na avaliacao do aluno.")

	mediaAceita := int(6)
	media := (n1 + n2) / 2

	if media >= mediaAceita {
		return true
	}
	return false

}

func ExecuteAvaliacaoAlunoAdiando() {
	fmt.Println(AvaliacaoAlunoAdiando(10, 10))
}
