package main

import (
	"fmt"
	"reflect"
)

func main() {
	evenOrOdd()
	//fmt.Println("Hello Word")
	//inferenciaDeTipo()

	//fmt.Println("Informe um número para o cálculo do fatorial")
	//var numero int
	//fmt.Scanf("%d", &numero)

	//fmt.Println(fatorial(numero))
}

func inferenciaDeTipo() {
	versao := 1.1
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("Inferindo o tipo da variável float. Há dois tipos de float em GO float32 e float64")
	fmt.Println("O tipo da variável idade é", reflect.TypeOf(versao))
}

func fatorial(numero int) int {
	if numero > 0 {
		resultado := numero * fatorial(numero-1)
		return resultado
	}
	return 1
}

func evenOrOdd() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Println(n, " é par")
		} else {
			fmt.Println(n, " é impar")
		}
	}
}
