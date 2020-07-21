package accounts

import "fmt"

type Account struct {
	name          string
	agency        int
	saldo         Saldo
	limit         float64
	lastOperation string
}

func (account *Account) deposit(value float64) string {
	message := account.saldo.deposit(value)
	fmt.Println(message)
	return message
}

func (account *Account) draw(value float64) string {
	message := account.saldo.draw(value)
	fmt.Println(message)
	return message
}

func (account *Account) transfer(value float64, receiver *Account) string {

	if account.saldo.authorizeRemoval(value) {

		account.saldo.draw(value)
		receiver.saldo.deposit(value)
		fmt.Println("Transferência concluída")

	} else {

		fmt.Println("Transferência impossibilitada")

	}

	return "---------------------------------------"
}
