package dialogues

import "strings"

//Dialog is ...
type Dialog struct {
	questions []string //TODO Criar um tipo question
	answers   []string //TODO Criar um tipo answer
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

func (d *Dialog) AmountOfQuestions() int {
	return len(d.questions)
}

//answer is ...
func (d *Dialog) answer(index int) (string, int) {
	if len(d.answers) < index {
		return "Answers not found", -1
	}
	return d.answers[index], index
}

func (d *Dialog) IsValidAnswer(userAnswer string, index int) bool {
	expectedAnswer, index := d.answer(index)
	if index == -1 {
		return false
	}
	return equals(userAnswer, expectedAnswer)
}

func equals(userAnswer string, expectedAnswer string) bool {
	return fmtAnswer(userAnswer) == fmtAnswer(expectedAnswer)
}

func fmtAnswer(answer string) string {
	return strings.ToUpper(strings.ReplaceAll(answer, " ", ""))
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
	return Dialog{questions: strQuestions, answers: strAnswers}
}
