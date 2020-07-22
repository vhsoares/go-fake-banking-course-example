package main

import (
	"fmt"

	"./accounts"
)

// classe para exibição do funcionamento do pacote de Conta
func main() {
	firstAccount := accounts.Account{"Vitor", 1255, accounts.Saldo{3000., 600.}, "None"}

	firstAccount.Deposit(200)

	fmt.Println(firstAccount)
}
