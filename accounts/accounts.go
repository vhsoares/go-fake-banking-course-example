package accounts

import "fmt"

type Account struct {
	Name          string
	Agency        int
	Saldo         Saldo
	LastOperation string
}

func (account *Account) Deposit(value float64) string {
	message := account.Saldo.Deposit(value)
	fmt.Println(message)
	return message
}

func (account *Account) Draw(value float64) string {
	message := account.Saldo.Draw(value)
	fmt.Println(message)
	return message
}

func (account *Account) Transfer(value float64, receiver *Account) string {

	if account.Saldo.authorizeRemoval(value) {

		account.Saldo.Draw(value)
		receiver.Saldo.Deposit(value)
		fmt.Println("Transferência concluída")

	} else {

		fmt.Println("Transferência impossibilitada")

	}

	return "---------------------------------------"
}
