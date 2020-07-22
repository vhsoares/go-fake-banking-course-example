package main

import "fmt"

type ContaCorrente struct {
	titular       string
	saldo         float64
	limite        float64
	numeroAgencia int
	numeroConta   int
}

func main() {

	curAcc := ContaCorrente{"Vitor", 8960.20, 5000, 1585, 115852201}
	fmt.Println(curAcc)

	// passando por cópia
	curAcc = curAcc.draw(1000)
	fmt.Println(curAcc)

	// passando por Referência
	curAcc.drawReference(1000)
	fmt.Println(curAcc)

}

func (conta ContaCorrente) draw(saque float64) ContaCorrente {

	if saque < 0 {
		return conta
	}

	if conta.saldo+conta.limite-saque > 0 {
		conta.saldo = conta.saldo - saque
	}

	return conta
}

func (conta *ContaCorrente) drawReference(saque float64) {

	if saque < 0 {
		return
	}

	if conta.saldo+conta.limite-saque > 0 {
		conta.saldo = conta.saldo - saque
	}
}
