package contas

import (
	"start-go/clientes"
)

//ContaCorrente structure
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo         float64
}

//Sacar withdraw money
func (c *ContaCorrente) Sacar(valorSaque float64) string {
	if valorSaque > 0 {
		if valorSaque <= c.saldo {
			c.saldo -= valorSaque
			return "Saque resalizado com sucesso"
		} else {
			return "Saldo insuficiente"
		}
	} else {
		return "Verifique o valor informado"
	}

}

//Depositar Add money
func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado", c.saldo
	} else {
		return "Valor inválido", c.saldo
	}
}

// Transferir Money
func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) string {
	if c.saldo > valorTransferencia && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return "Transferencia efetuada com sucesso"
	} else {
		return "não foi possivel transferir"
	}
}

// ObterSaldo get value
func (c *ContaCorrente) ObterSaldo () float64 {
	return c.saldo
}
