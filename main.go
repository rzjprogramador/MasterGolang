package main

import (
	"fmt"

	"github.com/MasterGolang/conversoes"
	"github.com/MasterGolang/escritas"
	"github.com/MasterGolang/loops"
	"github.com/MasterGolang/modules/math"
	"github.com/MasterGolang/objetos"
	"github.com/MasterGolang/operadores"
	"github.com/MasterGolang/tiposDeDados"
	"github.com/MasterGolang/variaveis"
)

func main() {
	// print("Hello Word - Main")
	variaveis.Declaracoes()
	conversoes.StringParaNumero()
	escritas.PrintConsole()
	loops.Incrementar()
	loops.PrintarItems()
	loops.VariarEmString()
	loops.VariarPorObjetoMap()
	fmt.Println(math.Soma(10, 20))
	objetos.StructModeloEstrutura()
	objetos.HerancaEstrutura()
	objetos.ObjRigidoMap()
	operadores.Logicos()
	operadores.Unarios()
	tiposDeDados.Erro()
	tiposDeDados.NumFloat()
	tiposDeDados.NumInteiro()
	tiposDeDados.ValoresIniciais()
}
