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
	dialog := new(dialogues.Dialog).Builder(lines)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < dialog.AmountOfQuestions(); i++ {

		question, index := dialog.Question(i)
		fmt.Println(question)

		scanner.Scan()
		userAnswer := scanner.Text()
		fmt.Println(dialog.IsValidAnswer(userAnswer, index))

	}
}

func shuffle(p []string) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(len(p) - 1)
}
