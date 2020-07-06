package contas

type Conta interface {
	Pagador
	Transferidor
	Depositador
}

type Pagador interface {
	Sacar(valor float64) string
}

type Transferidor interface {
	Transferir(valor float64, contaDestino Conta) bool
}

type Depositador interface {
	Depositar(valor float64) (string, float64)
}

func PagarBoleto(contaOrigem Conta, valorDoBoleto float64) {
	contaOrigem.Sacar(valorDoBoleto)
}
