package main

import "fmt"

//ContaCorrente structure
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
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
func (c *ContaCorrente) Depositar(valorDeposito float64) string {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado"
	} else {
		return "Valor inválido"
	}
}

func main() {
	contaMarco := ContaCorrente{}
	contaMarco.titular = "Marco Barroso"
	contaMarco.saldo = 1445.5

	fmt.Println(contaMarco.Sacar(1000.))
	fmt.Println(contaMarco.Depositar(245.))

	fmt.Println("Saldo:", contaMarco.saldo)
}
