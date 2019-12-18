package main

import "fmt"

//ContaCorrente sdfsdf
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// Example 1
	// var contaMarco ContaCorrente = ContaCorrente{}

	//Example 2
	contaMarco := ContaCorrente{titular: "Marco Barroso", numeroAgencia: 7857, numeroConta: 065473, saldo: 1200} //atribuição curta

	//Example 3
	contaTalita := ContaCorrente{"Talita Queiroz", 7887, 1234567, 120.40}

	//Example 4
	var contaTaina *ContaCorrente
	contaTaina = new(ContaCorrente)
	contaTaina.titular = "Tainá Queiroz"

	fmt.Println(contaMarco)
	fmt.Println(contaTalita)
	fmt.Println(*contaTaina)

	//Comparação 1
	contaTeste1 := ContaCorrente{"Talita Queiroz", 7887, 1234567, 120.40}
	contaTeste2 := ContaCorrente{"Talita Queiroz", 7887, 1234567, 120.40}
	fmt.Println(contaTeste1 == contaTeste2)

	//Comparação 2
	contaTeste3 := ContaCorrente{titular: "Marco Barroso", numeroAgencia: 7857, numeroConta: 065473, saldo: 1200} //atribuição curta
	contaTeste4 := ContaCorrente{titular: "Marco Barroso", numeroAgencia: 7857, numeroConta: 065473, saldo: 1200} //atribuição curta
	fmt.Println(contaTeste3 == contaTeste4)

	//Comparação 3
	var contaTest5 *ContaCorrente
	contaTest5 = new(ContaCorrente)
	contaTest5.titular = "Tainá Queiroz"

	var contaTest6 *ContaCorrente
	contaTest6 = new(ContaCorrente)
	contaTest6.titular = "Tainá Queiroz"
	fmt.Println(*contaTest5 == *contaTest6)

}
