package main

import (
	"fmt"
	"start-go/clientes"
	"start-go/contas"
)

//PagarBoleto Pay
func PagarBoleto(conta VerificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

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

	clienteTaina := clientes.Titular{
		"Tainá Queiroz",
		"457.111.657-45",
		"Infermeira",
	}

	//Create account
	contaMarco := contas.ContaCorrente{
		Titular:       clienteMarco,
		NumeroAgencia: 1788,
		NumeroConta:   12354,
	}
	contaTalita := contas.ContaCorrente{
		Titular:       clienteTalita,
		NumeroAgencia: 1754,
		NumeroConta:   654321,
	}

	contaTaina := contas.ContaPoupanca{
		Titular:       clienteTaina,
		NumeroAgencia: 1001,
		NumeroConta:   000454,
		Operacao:      78,
	}
	contaTaina.Depositar(150)
	fmt.Println(contaTaina)

	contaMarco.Depositar(1500)
	contaTalita.Depositar(1000)

	//Account operations
	fmt.Println(contaMarco.Sacar(100.))
	status, value := contaMarco.Depositar(245.)
	fmt.Println(status, value)

	//Example send money
	fmt.Println(contaMarco.Transferir(500, &contaTalita))
	fmt.Println("Saldo Marco:", contaMarco.ObterSaldo())
	fmt.Println("Saldo Talita:", contaTalita.ObterSaldo())

	//Pay
	PagarBoleto(&contaMarco, 145)
	fmt.Println("Saldo Marco:", contaMarco.ObterSaldo())

}
