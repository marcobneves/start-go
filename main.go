package main

import (
	"fmt"
	"start-go/contas"
)

func main() {
	//Create account
	contaMarco := contas.ContaCorrente{Titular: "Marco Barroso", Saldo: 1500}
	contaTalita := contas.ContaCorrente{Titular: "Talita Queiroz", Saldo: 1200}

	// Account operations
	fmt.Println(contaMarco.Sacar(100.))
	status, value := contaMarco.Depositar(245.)
	fmt.Println(status, value)

	// Example send money
	fmt.Println(contaMarco.Transferir(500, &contaTalita))
	fmt.Println("Saldo Marco:", contaMarco.Saldo)
	fmt.Println("Saldo Talita:", contaTalita.Saldo)

}
