package clientes

type Titular struct {
	Nome      string
	Cpf       string
	Profissao string
}

//Criar teste
func Criar(nome string, cpf string, profissao string) Titular {
	cliente := Titular{
		Nome:      nome,
		Cpf:       cpf,
		Profissao: profissao,
	}
	return cliente
}
