package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*

	Formas de declaração:

	A forma de declaração abaixo não é indicada há maiores chances de inversão dos atributos que causem bugs

		ricardo := pessoa{"Ludwig", "Ricardo"}

	De preferência ao formato a seguir

		ricardo := pessoa{nome: "Ricardo", sobrenome: "Ludwig"}

	Essa uma outra forma de atribuirmos valor

		var ricardo pessoa
		ricardo.nome = "Ricardo"
		ricardo.sobrenome = "Ludwig"

*/

type pessoa struct {
	//podemos usar tags de identificação para serialização em Json
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	contato
}

type contato struct {
	email string
	cep   int
}

func main() {

	fmt.Println("Struct: Estrutura de dados. Coleção de propriedades relacionadas")

	ricardo := pessoa{
		Nome:      "Ricardo",
		Sobrenome: "Ludwig",
		contato: contato{
			email: "ricardo@nowhere.com",
			cep:   32232,
		},
	}

	ricardo.print()

	out, err := json.MarshalIndent(ricardo, "", " ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	//precisa converter o array de bytes para string
	fmt.Println("JSON:", string(out))

}

//receiver function
func (p pessoa) print() {
	fmt.Printf("%+v\n", p)
}
