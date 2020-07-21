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

	// ponteiros com struct
	var cC *ContaCorrente
	fmt.Println(cC)
	cC = new(ContaCorrente)

	fmt.Println(cC)

}
