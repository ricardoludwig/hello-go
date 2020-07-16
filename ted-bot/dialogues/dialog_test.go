package dialogues

import "testing"

const howAreyou = "How are you?"
const imOk = "Im ok"

func TestSpligQuestionAndAnswesrs(t *testing.T) {

	dialog := splitQuestionAndAnswers(lines())

	if dialog.questions[0] != howAreyou {
		t.Error("Expected "+howAreyou+" , but got ", dialog.questions[0])
	}

	if dialog.answers[0] != imOk {
		t.Error("Expected "+imOk+", but got ", dialog.answers[0])
	}
}

func TestQuestion(t *testing.T) {

	dialog := builder()
	question, index := dialog.Question(0)

	if question != howAreyou {
		t.Error("Expected question "+howAreyou+" but got ", question)
	}

	if index != 0 {
		t.Error("Expected index 0 "+"but got ", index)
	}

}

func TestAnswer(t *testing.T) {

	dialog := builder()
	answer, index := dialog.answer(0)

	if answer != imOk {
		t.Error("Expected answer "+imOk+" but got ", answer)
	}

	if index != 0 {
		t.Error("Expected index 0 "+"but got ", index)
	}

}

func TestIsValidAnswer(t *testing.T) {

	dialog := builder()
	expectedAnswer, index := dialog.answer(0)
	userAnswer := "Im ok"

	if dialog.IsValidAnswer(userAnswer, index) == false {
		t.Error("Answer", userAnswer, " is not valid. Expected ", expectedAnswer)
	}
}

func builder() Dialog {
	return new(Dialog).Builder(lines())
}

func lines() []string {
	return []string{howAreyou, imOk}
}
