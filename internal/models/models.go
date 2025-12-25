package models

type Question struct {
	Name    string
	Code    string
	Answers []AnswerVariant
}
type AnswerVariant struct {
	Name    string
	Code    string
	IsRight bool
}

type Answer struct {
	Name string
	Code string
}

type ResponseQuestion struct {
	Name    string
	Code    string
	Answers []Answer
}

type UserAnswer struct {
	QuestionCode string   `json:"question_code"`
	AnswerCodes  []string `json:"answerCodes"`
}

type CheckResult struct {
	QuestionCode string   `json:"question_code"`
	UserAnswer   []string `json:"userAnswerCode"`
	RightAnswer  []string `json:"rightAnswerCode"`
	IsCorrect    bool     `json:"isCorrect"`
}

type NoiseSource struct {
	Name string
	Code string
}

type Quest struct {
	Noise    []NoiseSource
	Code     string
	ImagePNG string
}

type StatusThema string

const (
	StatusThemaOpen      StatusThema = `open`
	StatusThemaCurrent   StatusThema = `current`
	StatusThemaCompleted StatusThema = `completed`
)

type ThemaFullInfo struct {
	Info      Thema
	InfoBlock string
	Questions []Question
}
type Thema struct {
	Name   string
	Code   string
	Status StatusThema
}

type Course struct {
	Themas      []Thema
	Description string
}
