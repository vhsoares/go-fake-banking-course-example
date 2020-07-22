package main

import (
	"fmt"

	"./accounts"
)

// classe para exibição do funcionamento do pacote de Conta
func main() {
	firstAccount := accounts.Account{"Vitor", 1255, accounts.Saldo{3000., 600.}, "None"}
	secondAccount := accounts.Account{"Vitor", 1255, accounts.Saldo{3000., 600.}, "None"}

	accountVerification(firstAccount, secondAccount)

	firstAccount.Deposit(200)
	accountVerification(firstAccount, secondAccount)

	secondAccount.Draw(500)
	accountVerification(firstAccount, secondAccount)

	firstAccount.Transfer(1000, &secondAccount)
	accountVerification(firstAccount, secondAccount)
}

func accountVerification(firstAccount accounts.Account, secondAccount accounts.Account) {
	fmt.Println("Conta 1: ", firstAccount, "Conta 2: ", secondAccount)
}
