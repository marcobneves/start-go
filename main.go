package main

import (
	"fmt"
	"start-go/clientes"
	"start-go/contas"
)

func main() {

	//Create Client
	clienteMarco := clientes.Titular{
		Nome:      "Marco Barroso",
		Cpf:       "545.879.456-91",
		Profissao: "Desenvolvedor",
	}

	clienteTalita := clientes.Titular{
		"Talita Queiroz",
		"654.879.123-95",
		"Design",
	}

	//Create account
	contaMarco := contas.ContaCorrente{
		Titular:       clienteMarco,
		NumeroAgencia: 1788,
		NumeroConta:   12354,
		Saldo:         1500,
	}
	contaTalita := contas.ContaCorrente{
		Titular:       clienteTalita,
		NumeroAgencia: 1754,
		NumeroConta:   654321,
		Saldo:         100,
	}

	// // Account operations
	fmt.Println(contaMarco.Sacar(100.))
	status, value := contaMarco.Depositar(245.)
	fmt.Println(status, value)

	// // Example send money
	fmt.Println(contaMarco.Transferir(500, &contaTalita))
	fmt.Println("Saldo Marco:", contaMarco.Saldo)
	fmt.Println("Saldo Talita:", contaTalita.Saldo)

}
