package formas

import "fmt"

type Car struct {
	nome      string
	numeracao uint
}

var car1 = Car{"Fusca", 1000}

func useCreateCar(inputCar Car) Car {
	car := inputCar
	return car
}

// add metodo no prototipo da  struct de Car
func (c Car) andarProto() string {
	return c.nome + " esta andando"
}

func (c Car) showNumeracaoProto() uint {
	return c.numeracao
}

func CreateCar() {
	fmt.Println(useCreateCar(car1))
	fmt.Println(car1.andarProto())
	fmt.Println(car1.showNumeracaoProto())
}

/*

add_metodo_no_prototipo_da_estrutura:
	cria acao antes do nome (v Estrutura) deixe uma var que vai incorporar a estrutura semelhante a um this
	entao: atraves desta var vc tera acesso as props desta estrutura incorporada.
	e_assim: este metodo criado vai fazer parte desta estrutura como se fosse um emtodo da classe/struct no prototipo de todos novos objetos que serao criados.

*/
