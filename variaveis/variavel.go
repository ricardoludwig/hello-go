package main

import "fmt"

// Declaração de variáveis de escopo de pacote
var (
	//Endereco é um valor público
	Nome         string
	sobrenome    string
	idade        int
	maiorDeIdade bool
	salario      float64
	descricao    rune
)

func main() {

	fmt.Printf("Nome: %s\r\n", Nome)
	fmt.Printf("Sobrenome: %s\r\n", sobrenome)
	fmt.Printf("Tem maioridade? %v\r\n", maiorDeIdade)
	fmt.Printf("Idade: %d\r\n", idade)
	fmt.Printf("Salario:  %v\r\n", salario)

	//declaração de variável de escopo de função
	cidade := "São Paulo"
	fmt.Printf("Cidade: %s\r\n", cidade)

}
