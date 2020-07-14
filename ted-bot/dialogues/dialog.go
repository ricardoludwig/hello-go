package dialogues

//Dialog is ...
type Dialog struct {
	questions []string
	anwsers   []string
}

func (d *Dialog) Builder(lines []string) Dialog {
	if len(d.questions) == 0 {
		return splitQuestionAndAnswers(lines)
	}
	return Dialog{}
}

//Question is ...
func (d *Dialog) Question(index int) (string, int) {
	if len(d.questions) < index {
		return "Question not found", -1
	}
	return d.questions[index], index
}

func splitQuestionAndAnswers(lines []string) Dialog {
	var strQuestions []string
	var strAnswers []string
	for i, line := range lines {
		if i%2 == 0 {
			strQuestions = append(strQuestions, line)
		} else {
			strAnswers = append(strAnswers, line)
		}
	}
	return Dialog{questions: strQuestions, anwsers: strAnswers}
}
