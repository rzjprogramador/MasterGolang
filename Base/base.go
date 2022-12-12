package main

import (
	"github.com/MasterGolang/Base/agrupamento"
	"github.com/MasterGolang/Base/conversoes"
	"github.com/MasterGolang/Base/escritas"
	"github.com/MasterGolang/Base/fluxo"
	"github.com/MasterGolang/Base/loops"
	"github.com/MasterGolang/Base/objetos"
	"github.com/MasterGolang/Base/operadores"
	"github.com/MasterGolang/Base/tiposDeDados"
	"github.com/MasterGolang/Base/variaveis"
)

func main() {
	print("**** main do Base **** \n")
	variaveis.Declaracoes()
	conversoes.StringParaNumero()
	escritas.PrintConsole()
	loops.Incrementar()
	loops.PrintarItems()
	loops.VariarEmString()
	loops.VariarPorObjetoMap()
	objetos.StructModeloEstrutura()
	objetos.ComposicaoEstrutura()
	objetos.ObjRigidoMap()
	operadores.Logicos()
	operadores.Unarios()
	operadores.Relacionais()
	tiposDeDados.Erro()
	tiposDeDados.NumFloat()
	tiposDeDados.NumInteiro()
	tiposDeDados.ValoresIniciais()
	agrupamento.ArrayFixo()
	agrupamento.SliceArrayLivre()
	fluxo.SwitchCase()

}
