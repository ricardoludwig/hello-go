package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ricardoludwig/hello-go/ted-bot/ios"
)

func main() {

	dialog := ios.LoaderFile(ios.StringFileName, ios.StringPath)
	fmt.Println(dialog)

	for {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		origTexto := scanner.Text()

		fmt.Println(origTexto)

	}
}

func perguntaEReposta(pergunta string, respostas []string) string {

	for i, pg := range respostas {
		if pg == pergunta {
			return respostas[i]
		}
	}
	return "Errou!"
}

func separaPerguntaResposta(pr []string) ([]string, []string) {
	var perguntas []string
	var respostas []string
	for i, linha := range pr {
		if i%2 == 0 {
			perguntas = append(perguntas, linha)
		} else {
			respostas = append(respostas, linha)
		}
	}
	return perguntas, respostas
}

func shuffle(p []string) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(len(p) - 1)
}

func test() {

	dialog := ios.LoaderFile(ios.StringFileName, ios.StringPath)
	perguntas, respostas := separaPerguntaResposta(dialog)

	for {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		origTexto := scanner.Text()

		fmt.Println(origTexto)

		if "" == strings.ReplaceAll(strings.ToUpper(origTexto), " ", "") {
			fmt.Println("Não respondeu nada, né? Você não sabe?")
		} else {
			if strings.ReplaceAll(strings.ToUpper(origTexto), " ", "") == "OI" {
				fmt.Println(respostas[0])
			}
			i := shuffle(perguntas)
			fmt.Println(perguntas[i])
			if strings.ReplaceAll(strings.ToUpper(respostas[i]), " ", "") == strings.ReplaceAll(strings.ToUpper(origTexto), " ", "") {
				fmt.Println("Você acertou", respostas[i])
			} else {
				fmt.Println("Teste:", origTexto)
				fmt.Println("RespostA", respostas[i])
				fmt.Println("Não está certo, você pode estudar mais e tentar da próxima vez")
			}
		}
	}
}
