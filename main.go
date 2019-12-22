package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"start-go/clientes"

	_ "github.com/lib/pq"
)

//ConectaBanco Connect database
func ConectaBanco() *sql.DB {
	conexao := "postgres://postgres:Postgres2019!@postgres-db/start-go?sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

//PagarBoleto Pay
func PagarBoleto(conta VerificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

//Load all files html
var temp = template.Must(template.ParseGlob("templates/*.html"))

//Index show template
func Index(w http.ResponseWriter, r *http.Request) {
	db := ConectaBanco()
	ClientesBanco, err := db.Query("SELECT * FROM clientes")

	if err != nil {
		panic(err.Error())
	}
	Clientes := []clientes.Titular{}

	for ClientesBanco.Next() {
		var Id int
		var Nome, Cpf, Profissao string
		err := ClientesBanco.Scan(&Id, &Nome, &Cpf, &Profissao)
		fmt.Println("Nome", Nome)
		if err != nil {
			panic(err.Error())
		}
		NovoCliente := clientes.Criar(Nome, Cpf, Profissao)
		Clientes = append(Clientes, NovoCliente)
	}
	temp.ExecuteTemplate(w, "Index", Clientes)
	defer db.Close()
}

func main() {

	// sites := []string{"google", "yahooo", "Facebook", "G1"}
	// fmt.Println(sites)

	// for i := 1; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	// for i, site := range sites {
	// 	fmt.Println(i, "-", site)
	// }

	//Conect
	// db := ConectaBanco()
	// defer db.Close()

	// Listen route '/'
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)

	// //Create Client
	// clienteMarco := clientes.Titular{
	// 	Nome:      "Marco Barroso",
	// 	Cpf:       "545.879.456-91",
	// 	Profissao: "Desenvolvedor",
	// }

	// clienteTalita := clientes.Titular{
	// 	"Talita Queiroz",
	// 	"654.879.123-95",
	// 	"Design",
	// }

	// clienteTaina := clientes.Titular{
	// 	"Tainá Queiroz",
	// 	"457.111.657-45",
	// 	"Infermeira",
	// }

	// //Create account
	// contaMarco := contas.ContaCorrente{
	// 	Titular:       clienteMarco,
	// 	NumeroAgencia: 1788,
	// 	NumeroConta:   12354,
	// }
	// contaTalita := contas.ContaCorrente{
	// 	Titular:       clienteTalita,
	// 	NumeroAgencia: 1754,
	// 	NumeroConta:   654321,
	// }

	// contaTaina := contas.ContaPoupanca{
	// 	Titular:       clienteTaina,
	// 	NumeroAgencia: 1001,
	// 	NumeroConta:   000454,
	// 	Operacao:      78,
	// }
	// contaTaina.Depositar(150)
	// fmt.Println(contaTaina)

	// contaMarco.Depositar(1500)
	// contaTalita.Depositar(1000)

	// //Account operations
	// fmt.Println(contaMarco.Sacar(100.))
	// status, value := contaMarco.Depositar(245.)
	// fmt.Println(status, value)

	// //Example send money
	// fmt.Println(contaMarco.Transferir(500, &contaTalita))
	// fmt.Println("Saldo Marco:", contaMarco.ObterSaldo())
	// fmt.Println("Saldo Talita:", contaTalita.ObterSaldo())

	// //Pay
	// PagarBoleto(&contaMarco, 145)
	// fmt.Println("Saldo Marco:", contaMarco.ObterSaldo())

}
