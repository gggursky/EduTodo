package service

import (
	"github.com/Konstantin299/EduTodo.git/internal/models"
	"github.com/sirupsen/logrus"
)

type Service struct {
	log *logrus.Entry
}

func New(log *logrus.Logger) *Service {
	return &Service{
		log: log.WithField("module", "service"),
	}
}

func (s *Service) GetInfo() (string, error) {
	h := "hello"

	return h, nil
}

//func (s *Service) GetAnswersList() ([]string, error) {
//	return []string{answer1.Answer, answer2.Answer}, nil
//}

func (s *Service) GetAnswersList() ([]models.ResponseQuestion, error) {
	var result []models.ResponseQuestion

	for _, q := range questions {
		var tmp models.ResponseQuestion

		tmp.Name = q.Name
		tmp.Code = q.Code

		for _, a := range q.Answers {
			tmp.Answers = append(tmp.Answers, models.Answer{Name: a.Name, Code: a.Code})
		}

		result = append(result, tmp)
	}

	return result, nil
}

//func (s *Service) CheckAnswer([]models.ResponseQuestion) (string, error) {
//	//for _, q := range questions {
//	//	if q.Code == questionCode {
//	//		for _, a := range q.Answers {
//	//			if a.Code == respCode {
//	//				if a.IsRight {
//	//					return fmt.Sprintf("Верно: вопрос \"%s\", выбран ответ \"%s\" (код %s)", q.Name, a.Name, a.Code), nil
//	//				}
//	//				return fmt.Sprintf("Неверно: вопрос \"%s\", выбран ответ \"%s\" (код %s)", q.Name, a.Name, a.Code), nil
//	//			}
//	//		}
//	//		return "", fmt.Errorf("ответ с кодом %s для вопроса %s не найден", respCode, questionCode)
//	//	}
//	//}
//	//return "", fmt.Errorf("вопрос с кодом %s не найден", questionCode)
//	return "Ок", nil
//}

//func (s *Service) CheckAnswer(responses []models.UserAnswer) (string, error) {
//	correct := 0
//
//	for _, resp := range responses {
//		// ищем исходный вопрос
//		for _, q := range questions {
//			if q.Code == resp.QuestionCode {
//
//				// проверяем выбранный ответ
//				for _, ans := range q.Answers {
//					if ans.Code == resp.AnswerCode && ans.IsRight {
//						correct++
//					}
//				}
//			}
//		}
//	}
//
//	return fmt.Sprintf("Correct answers: %d/%d", correct, len(questions)), nil
//}

func sameSet(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// создаём карту для элементов a
	m := make(map[string]struct{})
	for _, x := range a {
		m[x] = struct{}{}
	}

	// проверяем, что каждый элемент b есть в карте
	for _, x := range b {
		if _, ok := m[x]; !ok {
			return false
		}
	}

	return true
}

func (s *Service) CheckAnswer(userAnswers []models.UserAnswer) ([]models.CheckResult, error) {
	var results []models.CheckResult

	for _, ua := range userAnswers {
		result := models.CheckResult{
			QuestionCode: ua.QuestionCode,
			UserAnswer:   ua.AnswerCodes,
			IsCorrect:    false,
		}

		for _, q := range questions {
			if q.Code != ua.QuestionCode {
				continue
			}

			var right []string
			for _, ans := range q.Answers {
				if ans.IsRight {
					right = append(right, ans.Code)
				}
			}

			result.RightAnswer = right
			result.IsCorrect = sameSet(ua.AnswerCodes, right)

			break
		}

		results = append(results, result)
	}

	return results, nil
}
