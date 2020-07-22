package accounts

type Saldo struct {
	Value float64
	Limit float64
}

func (saldo *Saldo) Deposit(value float64) string {

	if negativeValue(value) {
		return "Operação não autorizada: Valor Negativo"
	}

	saldo.Value = saldo.Value + value
	return "Depósito efetuado com sucesso"
}

func (saldo *Saldo) Draw(value float64) string {

	if negativeValue(value) {
		return "Operação não autorizada: Valor Negativo"
	}

	if saldo.authorizeRemoval(value) {
		saldo.Value = saldo.Value - value
		return "Operação Concluída"
	}

	return "Saldo Insuficiente"

}

func (saldo *Saldo) authorizeRemoval(value float64) bool {
	return saldo.Value+saldo.Limit < value
}

func negativeValue(value float64) bool {
	return value < 0
}
