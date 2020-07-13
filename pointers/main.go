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

Argumentos passados para funções são copiados e usados dentro dela, Go não altera o falor original.
Exceto quando o valor alterado for um item de um Slice.
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

	/* Não é necessário fazer a declaração dessa forma, como estamos definindo na
	declaração da função atualizaNome o *pessoa, go é capaz de fazer essas operações
	automaticamente

	ricardoPonteiro := &ricardo // & --> Obtem o endereço de memória do valor da variável
	ricardoPonteiro.atualizaNome("Ludwig")

	*/

	ricardo.atualizaNome("Ludwig")
	fmt.Println("Após a atualização")
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

//*pessoa --> É um tipo de descrição que indica que estamos trabalhando com ponteiros para aquele tipo (type)
func (ponteiroParaPessoa *pessoa) atualizaNome(novoPrimeiroNome string) {
	(*ponteiroParaPessoa).nome = novoPrimeiroNome // *ponteiroParaPessoa --> Obtem o valor do endereço da memória. Nesse caso * é um operador
}

/*
Resumindo:
 Transformar endereço de memória em um valor --> usar *address
 Transformar um valor em endereço de memória --> usar &value
*/
