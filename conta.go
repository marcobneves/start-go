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

func main() {
	contaMarco := ContaCorrente{}
	contaMarco.titular = "Marco Barroso"
	contaMarco.saldo = 1445.5
	resultado := contaMarco.Sacar(10000.)

	fmt.Println(resultado)
	fmt.Println("Saldo:", contaMarco.saldo)
}
