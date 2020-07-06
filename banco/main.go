package main

import (
	"fmt"

	"github.com/ricardoludwig/hello-go/banco/clientes"
	"github.com/ricardoludwig/hello-go/banco/contas"
)

func main() {

	contaDoRicardo := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Ricardo",
		CPF:       "123.456.789-01",
		Profissao: "Dev"},
		NumeroAgencia: 1234, NumeroConta: 567}

	contaDoRicardo.Depositar(2000)
	fmt.Println(contaDoRicardo)

	contas.PagarBoleto(&contaDoRicardo, 100)
	fmt.Println(contaDoRicardo.ObterSaldo())

	contaDoZezinho := contas.ContaPoupanca{Titular: clientes.Titular{
		Nome:      "Zezinho",
		CPF:       "339.503.890-09",
		Profissao: "Dev"},
		NumeroAgencia: 789, NumeroConta: 1036}
	contaDoZezinho.Depositar(5000)

	contas.PagarBoleto(&contaDoZezinho, 300)
	fmt.Println(contaDoZezinho.ObterSaldo())

	sucessoNaTransferencia := contaDoRicardo.Transferir(100, &contaDoZezinho)
	if sucessoNaTransferencia {
		fmt.Println("Tranferencia realizada com sucesso")
		fmt.Println("saldo conta Ricardo", contaDoRicardo.ObterSaldo())
		fmt.Println("saldo conta Zezinho", contaDoZezinho.ObterSaldo())
	} else {
		fmt.Println("Falha na tranferencia")
	}

}
