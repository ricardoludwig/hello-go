package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Oi como vai você? Me chamo Lynx. Qual o seu nome?")

	for {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		origTexto := scanner.Text()

		texto := strings.ReplaceAll(origTexto, " ", "")
		texto = strings.ToUpper(texto)

		switch texto {
		case "OI":
			fmt.Println("O que você está fazendo de bom?")
		case "NADA":
			fmt.Println("Não tem nada de bom para fazer?")
		case "JOGAR":
			fmt.Println("Vocẽ não pensa em fazer outra coisa além de jogar?")
		case "JOGANDO":
			fmt.Println("Você não estuda?")
		case "ESTUDO":
			fmt.Println("Não é o que sua mãe fala")
		case "ESTOUBEM":
			fmt.Println("Felipe, quanto é 2x2?")
		case "4":
			fmt.Println("Muito bem, você é muito inteligente, mas eu dúvido você acertar essa")
			fmt.Println("Quanto é a raiz quadrada de 4?")
		case "2":
			fmt.Println("Parabéns! Essa será muito difícil")
			fmt.Println("Quanto é 110 - 85? Eu não sei me ajude!! Preciso resolver esse mistério")
		case "25":
			fmt.Println("Quanto é 10 x 10? Essa você não vai acertar. Dúvido você acertar!")
		case "100":
			fmt.Println("Essa eu achei que você não iria acertar. Me fala qual é a capital do Brasil")
		case "BRASILIA":
			fmt.Println("Muito bem, e qual é o melhor time de futebol do mundo?")
		case "SAOPAULO":
			fmt.Println("Você tem razão, é o melhor time do mundo Tricampeão Mundial!!! Salve o tricolor Paulista, amado clube!!")
			fmt.Println("Felipe quantos vezes o São Paulo ganhou a libertadores e o Mundial?")
		case "3":
			fmt.Println("Por hoje é só Felipe")
			os.Exit(0)
		case "SANTOS":
			fmt.Println("HAHUAHAHUA HAUHAUHAUHAHUA HAHUAHUAHAHU conta outra piada!!! Por favor, essa foi muito boa")
			fmt.Println("Tenta outra vez, quem sabe você acerta")
		case "BARCELONA":
			fmt.Println("É um time muito bom, tem o Messi! Mas sabe qual time já foi campeão mundial ganhando do Barcelona?")
		case "FELIPE":
			fmt.Println("Bonito nome, eu aposto que seu tio tem um nome bonito também. Qual o nome dele?")
		case "RICARDO":
			fmt.Println("Qual o nome da sua mãe?")
		case "RENATA":
			fmt.Println("Que nome lindo, aposto que ela é muito bonita e gosta muito de você")
			fmt.Println("Podemos continuar conversando, responde Sim ou Não")
		case "SIM":
			fmt.Println("Ebaaaa! Que bom que vamos continuar eu estou gostando muito de conversar com você")
			fmt.Println("Vou te fazer uma pergunta? Quanto é 1+1?")
		case "NAO":
			fmt.Println("Que pena, gostei muito de conversar com você, estarei sempre aqui quando quiser conversar mais")
			fmt.Println("Tchau!!!")
			os.Exit(0)
		case "TCHAU":
			fmt.Println("Tchau, foi muito bom conversar com você")
			os.Exit(0)
		default:
			fmt.Println("Desculpe Felipe, eu ainda não sou muito inteligente. Mas eu sei quanto é 2x2. Você sabe?")
		}
	}
}
