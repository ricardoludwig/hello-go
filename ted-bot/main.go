package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ricardoludwig/hello-go/ted-bot/dialogues"
	"github.com/ricardoludwig/hello-go/ted-bot/ios"
)

func main() {

	lines := ios.LoaderFile(ios.StringFileName, ios.StringPath)
	fmt.Println(lines)

	dialog := new(dialogues.Dialog).Builder(lines)
	fmt.Println(dialog)

	for {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		origTexto := scanner.Text()

		fmt.Println(origTexto)

	}
}

func shuffle(p []string) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(len(p) - 1)
}
