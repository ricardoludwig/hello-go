package dialogues

type Dialog struct {
	language  string
	questions []question
}

type question struct {
	strQuestion string
	answers     []string
}

func splitQuestionAndAnswers(lines []string) ([]string, []string) {
	var strQuestions []string
	var strAnswers []string
	for i, line := range lines {
		if i%2 == 0 {
			strQuestions = append(strQuestions, line)
		} else {
			strAnswers = append(strAnswers, line)
		}
	}
	return strQuestions, strAnswers
}
