package accounts

type Saldo struct {
	value float64
	limit float64
}

func (saldo *Saldo) deposit(value float64) string {

	if positiveValue(value) {
		return "Operação não autorizada: Valor Negativo"
	}

	saldo.value = saldo.value + value
	return "Depósito efetuado com sucesso"
}

func (saldo *Saldo) draw(value float64) string {

	if positiveValue(value) {
		return "Operação não autorizada: Valor Negativo"
	}

	if saldo.authorizeRemoval(value) {
		saldo.value = saldo.value - value
		return "Operação Concluída"
	}

	return "Saldo Insuficiente"

}

func (saldo *Saldo) authorizeRemoval(value float64) bool {
	return saldo.value+saldo.limit < value
}

func positiveValue(value float64) bool {
	return value > 0
}
