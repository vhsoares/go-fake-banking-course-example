package main

import "fmt"

type ContaCorrente struct {
	saldo float64
	nome  string
}

func main() {
	var conta *ContaCorrente
	conta = new(ContaCorrente)
	conta.nome = "Vitor"
	conta.saldo = 500.
	conta.deposit(100)
	fmt.Println(conta.saldo)

	transferencia()
}

func transferencia() {
	contaTransfere := ContaCorrente{1000., "Vítor"}
	contaRecebe := ContaCorrente{1000., "Vítor"}

	status := transfer(&contaTransfere, &contaRecebe, 250)
	fmt.Println(status)
	fmt.Println(contaTransfere)
	fmt.Println(contaRecebe)
}

func transfer(contaTransfere *ContaCorrente, contaRecebe *ContaCorrente, valorTransferencia float64) string {
	if valorTransferencia < contaTransfere.saldo && valorTransferencia > 0 {
		contaTransfere.saldo -= valorTransferencia
		contaRecebe.saldo = contaRecebe.saldo + valorTransferencia
		return "Valor Transferido"
	}
	return "Não foi possível Transferir"
}

func (conta *ContaCorrente) deposit(valor float64) {

	if valor > 0 {
		conta.saldo = conta.saldo + valor
	} else {
		fmt.Println("Não pode adicionar valor negativo")
	}

}
