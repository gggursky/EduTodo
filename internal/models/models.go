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
	QuestionCode string   `json:"questionCode"`
	UserAnswer   []string `json:"userAnswerCode"`
	RightAnswer  []string `json:"rightAnswerCode"`
	IsCorrect    bool     `json:"isCorrect"`
}
