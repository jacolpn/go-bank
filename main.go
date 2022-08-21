package main

import (
	"fmt"
	"go-bank/contas"
)

func main() {
	// <------------------------>
	// var titular string = "Jackson"
	// var numeroAgencia int = 589
	// var numeroConta int = 123456
	// var saldo float64 = 125.50
	// fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	// var contaDoGuilherme ContaCorrente = ContaCorrente{
	// 	titular:       "jackson",
	// 	numeroAgencia: 12,
	// 	numeroConta:   1,
	// 	saldo:         1,
	// }

	// contaDoGuilherme := ContaCorrente{
	// 	titular:       "jackson",
	// 	numeroAgencia: 12,
	// 	numeroConta:   1,
	// 	saldo:         1,
	// }

	// contaDoGuilherme := ContaCorrente{"jackson", 12, 1, 1}

	// fmt.Println("Conta do Guilherme: ", contaDoGuilherme)

	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.titular = "Cris"

	// fmt.Println("Conta da Cris: ", contaDaCris)
	// fmt.Println("Conta da Cris (ponteiro): ", *contaDaCris)
	// <------------||------------>

	// <------------------------>
	// contaDaSilvia := ContaCorrente{}
	// contaDaSilvia.titular = "Silvia"
	// contaDaSilvia.saldo = 500

	// fmt.Println(contaDaSilvia)

	// valorDoSaque := 100
	// // contaDaSilvia.saldo = contaDaSilvia.saldo - float64(valorDoSaque)

	// fmt.Println(contaDaSilvia.saldo)
	// fmt.Println(contaDaSilvia.Sacar(float64(valorDoSaque)))
	// fmt.Println(contaDaSilvia.saldo)

	// status, valor := contaDaSilvia.Depositar(100)
	// // fmt.Println(contaDaSilvia.Depositar(100))
	// fmt.Println(status, valor)
	// fmt.Println(contaDaSilvia.saldo)
	// <------------||------------>

	// <------------------------>
	// contaDaSilvia := c.ContaCorrente{Titular: "Silvia", Saldo: 300}
	// contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	// status := contaDaSilvia.Transferir(200, &contaDoGustavo)

	// fmt.Println(status)
	// fmt.Println(contaDaSilvia)
	// fmt.Println(contaDoGustavo)
	// <------------||------------>

	// <------------------------>
	// contaDoJackson := contas.ContaCorrente{Titular: clientes.Titular{
	// 	Nome: "Jackson", CPF: "000.000.000-00", Profissao: "Developer"},
	// 	NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}

	// clienteJackson := clientes.Titular{"Jackson", "000.000.000-00", "Developer"}
	// contaJackson := contas.ContaCorrente{clienteJackson, 123, 123456, 100}

	// fmt.Println(clienteJackson, contaJackson)
	// <------------||------------>

	// <------------------------>
	// contaExemplo := contas.ContaCorrente{}
	// contaExemplo.Depositar(-100)

	// fmt.Println(contaExemplo, contaExemplo.ObterSaldo())
	// <------------||------------>

	contaDaAngelica := contas.ContaPoupanca{}
	contaDaAngelica.Depositar(100)
	PagarBoleto(&contaDaAngelica, 60)

	fmt.Println(contaDaAngelica)

	contaDaAngelica2 := contas.ContaCorrente{}
	contaDaAngelica2.Depositar(100)
	PagarBoleto(&contaDaAngelica2, 60)

	fmt.Println(contaDaAngelica2)
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
