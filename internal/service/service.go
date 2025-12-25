package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Konstantin299/EduTodo.git/internal/models"
	"github.com/sirupsen/logrus"
)

type Service struct {
	log    *logrus.Entry
	Themas map[string]models.ThemaFullInfo
}

func New(log *logrus.Logger) *Service {
	return &Service{
		log: log.WithField("module", "service"),
		Themas: map[string]models.ThemaFullInfo{
			"code1": thema1,
			"code2": thema2,
			"code3": thema3,
		},
	}
}

func (s *Service) GetInfo() (string, error) {
	h := "hello"

	return h, nil
}

//func (s *Service) GetAnswersList() ([]string, error) {
//	return []string{answer1.Answer, answer2.Answer}, nil
//}

//func (s *Service) GetAnswersList() ([]models.ResponseQuestion, error) {
//	var result []models.ResponseQuestion
//
//	for _, q := range questions {
//		var tmp models.ResponseQuestion
//
//		tmp.Name = q.Name
//		tmp.Code = q.Code
//
//		for _, a := range q.Answers {
//			tmp.Answers = append(tmp.Answers, models.Answer{Name: a.Name, Code: a.Code})
//		}
//
//		result = append(result, tmp)
//	}
//
//	return result, nil
//}

//func (s *Service) CheckAnswer([]models.ResponseQuestion) (string, error) {
//	for _, q := range questions {
//		if q.Code == questionCode {
//			for _, a := range q.Answers {
//				if a.Code == respCode {
//					if a.IsRight {
//						return fmt.Sprintf("Верно: вопрос \"%s\", выбран ответ \"%s\" (код %s)", q.Name, a.Name, a.Code), nil
//					}
//					return fmt.Sprintf("Неверно: вопрос \"%s\", выбран ответ \"%s\" (код %s)", q.Name, a.Name, a.Code), nil
//				}
//			}
//			return "", fmt.Errorf("ответ с кодом %s для вопроса %s не найден", respCode, questionCode)
//		}
//	}
//	return "", fmt.Errorf("вопрос с кодом %s не найден", questionCode)
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
	results := make([]models.CheckResult, 0, len(userAnswers))

	for _, ua := range userAnswers {

		var found bool

		for _, thema := range s.Themas {
			for _, q := range thema.Questions {
				if q.Code != ua.QuestionCode {
					continue
				}

				var right []string
				for _, a := range q.Answers {
					if a.IsRight {
						right = append(right, a.Code)
					}
				}

				results = append(results, models.CheckResult{
					QuestionCode: q.Code,
					UserAnswer:   ua.AnswerCodes,
					RightAnswer:  right,
					IsCorrect:    sameSet(ua.AnswerCodes, right),
				})

				found = true
				break
			}
			if found {
				break
			}
		}
	}

	return results, nil
}

//func (s *Service) GetQuest() (models.Quest, error) {
//	rand.Seed(time.Now().UnixNano())
//
//	idx := rand.Intn(len(quests)) // случайный индекс
//	return quests[idx], nil
//}
//

func (s *Service) GetThemas() ([]models.Thema, error) {
	return course, nil
}

func (s *Service) GetThemaFullInfo(code string) (models.ThemaFullInfo, error) {
	info, ok := s.Themas[code]
	if !ok {
		return models.ThemaFullInfo{}, fmt.Errorf("тема не найдена")
	}
	return info, nil
}

func (s *Service) GetQuestionsList(code string) []models.ResponseQuestion {
	thema, ok := s.Themas[code]
	if !ok {
		return nil
	}

	// ⚠️ важно инициализировать seed
	rand.Seed(time.Now().UnixNano())

	// копируем слайс, чтобы не портить исходный порядок
	questions := make([]models.Question, len(thema.Questions))
	copy(questions, thema.Questions)

	// перемешиваем
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	result := make([]models.ResponseQuestion, 0, len(questions))

	for _, q := range questions {

		// копируем ответы
		answersVariants := make([]models.AnswerVariant, len(q.Answers))
		copy(answersVariants, q.Answers)

		// перемешиваем
		rand.Shuffle(len(answersVariants), func(i, j int) {
			answersVariants[i], answersVariants[j] = answersVariants[j], answersVariants[i]
		})

		answers := make([]models.Answer, 0, len(answersVariants))
		for _, a := range answersVariants {
			answers = append(answers, models.Answer{
				Name: a.Name,
				Code: a.Code,
			})
		}

		result = append(result, models.ResponseQuestion{
			Name:    q.Name,
			Code:    q.Code,
			Answers: answers,
		})
	}
	return result
}
