package contas

import (
	"start-go/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

//Sacar withdraw money
func (c *ContaPoupanca) Sacar(valorSaque float64) string {
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
func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado", c.saldo
	} else {
		return "Valor inválido", c.saldo
	}
}

// ObterSaldo get value
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
