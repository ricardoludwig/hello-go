package main

import "fmt"

/*******
Tipos de desclarações

//Não é uma boa se inverter a declaração no struts vai gerar um bug
ricardo := pessoa{"Ricardo", "Ludwig"}

ricardo := pessoa{nome: "Ricardo", sobrenome: "Ludwig"}

var ricardo pessoa
ricardo.nome = "Ricardo"
ricardo.sobrenome = "Ludwig"

*******/

type contato struct {
	email string
	cep   int
}

type pessoa struct {
	nome      string
	sobrenome string
	contato
}

func main() {

	fmt.Println("Struct: Estrutura de dados. Coleção de propriedades relacionadas")

	ricardo := pessoa{
		nome:      "Ricardo",
		sobrenome: "Ludwig",
		contato: contato{
			email: "ricardo@nowhere.com",
			cep:   32232,
		},
	}

	ricardo.print()

	ricardoPonteiro := &ricardo
	ricardoPonteiro.atualizaNome("Ludwig")

	ricardo.print()
}

//Receiver Functions

func (p pessoa) print() {
	fmt.Printf("%+v", p)
}

//Passagem por valor (faz uma cópia da Struct em outro endereço de memória)
//Dessas forma não funciona, não vai atualizar
func (p pessoa) updateName(newFirstName string) {
	p.nome = newFirstName
}

func (ponteiroParaPessoa *pessoa) atualizaNome(novoPrimeiroNome string) {
	(*ponteiroParaPessoa).nome = novoPrimeiroNome
}