package main

import "fmt"

/*
	Exceto por algumas exceções em Go a passagem de argumentos é feito por cópia, ou seja, Go não altera o valor original.

	Resumo:

	Transformar endereço de memória em um valor --> usar *
 	Transformar um valor em endereço de memória --> usar &

*/

type pessoa struct {
	nome      string
	sobrenome string
	contato
}

type contato struct {
	email string
	cep   int
}

func main() {

	fmt.Println("Ponteiros")

	ricardo := pessoa{
		nome:      "Ricardo",
		sobrenome: "Ludwig",
		contato: contato{
			email: "ricardo@nowhere.com",
			cep:   32232,
		},
	}

	ricardo.updateNameFail("Qualquer")
	fmt.Println("Nome não atualizado")
	ricardo.print()

	fmt.Println("Nome atualizado")
	ricardo.updateName("Ludwig")
	ricardo.print()
}

// receiver function
func (p pessoa) print() {
	fmt.Printf("%+v\n", p)
}

// passagem por valor (faz uma cópia da Struct em outro endereço de memória);. não atualizará o nome
func (p pessoa) updateNameFail(newFirstName string) {
	p.nome = newFirstName
}

// *pessoa, é um tipo de descrição que indica que estamos trabalhando com ponteiros para aquele type
func (ponteiroParaPessoa *pessoa) updateName(novoPrimeiroNome string) {
	(*ponteiroParaPessoa).nome = novoPrimeiroNome // *ponteiroParaPessoa --> Obtem o valor do endereço da memória. Nesse caso * é um operador
}
