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

func main() {
	//Create account
	contaMarco := ContaCorrente{titular: "Marco Barroso", saldo: 1500}
	contaTalita := ContaCorrente{titular: "Talita Queiroz", saldo: 1200}

	// Account operations
	fmt.Println(contaMarco.Sacar(100.))
	status, value := contaMarco.Depositar(245.)
	fmt.Println(status, value)

	// Example send money
	fmt.Println(contaMarco.Transferir(500, &contaTalita))
	fmt.Println("Saldo Marco:", contaMarco.saldo)
	fmt.Println("Saldo Talita:", contaTalita.saldo)

}
