package rest

import (
	"net/http"

	"github.com/Konstantin299/EduTodo.git/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetText(context *gin.Context) {
	text, _ := s.service.GetInfo()

	context.IndentedJSON(
		http.StatusOK,
		text,
	)
}

// GetAnswers возвращает список вопросов
// @Summary возвращает список вопросов с вариантами ответов
// @Description возвращает список вопросов с вариантами ответов
// @Tags тесты
// @Accept json
// @Produce json
// @Success 201 {object} []models.ResponseQuestion
// @Router /answers [GET]
func (s *Server) GetAnswers(context *gin.Context) {
	answersList, _ := s.service.GetAnswersList()

	context.IndentedJSON(
		http.StatusOK,
		answersList,
	)

}

//func (s *Server) Check(context *gin.Context) {
//	type questions []models.ResponseQuestion
//
//	var request questions
//
//	_ = context.BindJSON(&request)
//
//	resp, _ := s.service.CheckAnswer(request)
//
//	context.IndentedJSON(
//		http.StatusOK,
//		resp,
//	)
//}

// Check проверяет ответы пользователя
// @Summary Проверка ответов пользователя
// @Description Принимает список выбранных пользователем ответов (questionCode + answerCodes) и возвращает результат проверки
// @Tags тесты
// @Accept json
// @Produce json
// @Param request body []models.UserAnswer true "Ответы пользователя"
// @Success 200 {array} models.CheckResult
// @Failure 400 {object} map[string]string "Ошибка запроса"
// @Failure 500 {object} map[string]string "Внутренняя ошибка"
// @Router /check [POST]
func (s *Server) Check(context *gin.Context) {
	var request []models.UserAnswer

	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.service.CheckAnswer(request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"result": result})
}

// GetQuest возвращает случайный вопрос
// @Summary Случайный вопрос
// @Description Возвращает случайно выбранный вопрос из базы
// @Tags тесты
// @Produce json
// @Success 200 {object} models.Quest
// @Router /quest [GET]
func (s *Server) GetQuest(context *gin.Context) {
	answersList, _ := s.service.GetQuest()

	context.IndentedJSON(
		http.StatusOK,
		answersList,
	)
}

// GetThemas возвращает список тем
// @Summary Список тем
// @Description Возвращает краткую информацию по всем темам курса
// @Tags темы
// @Produce json
// @Success 200 {array} models.Thema
// @Router /themas [GET]
func (s *Server) GetThemas(context *gin.Context) {
	answersList, _ := s.service.GetThemas()

	context.IndentedJSON(
		http.StatusOK,
		answersList,
	)

}

// GetThemaFullInfo возвращает текстовое описание темы
// @Summary Полная информация о теме
// @Description Возвращает текстовый блок (InfoBlock) по теме
// @Tags темы
// @Produce text/plain
// @Param code path string true "Код темы"
// @Success 200 {string} string "Текст темы"
// @Failure 404 {string} string "Тема не найдена"
// @Router /thema/{code} [GET]
func (s *Server) GetThemaFullInfo(c *gin.Context) {
	code := c.Param("code")

	info, err := s.service.GetThemaFullInfo(code)
	if err != nil {
		c.String(http.StatusNotFound, "Тема не найдена")
		return
	}

	// Отдаём чистый текст
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(info.InfoBlock))
}

// GetQuestionsList возвращает список вопросов по теме
// @Summary Список вопросов по теме
// @Description Возвращает список вопросов для указанной темы без вариантов ответов
// @Tags тесты
// @Produce json
// @Param code path string true "Код темы"
// @Success 200 {array} models.Question
// @Failure 404 {object} map[string]string "Вопросы для темы не найдены"
// @Router /thema/{code}/questions [GET]
func (s *Server) GetQuestionsList(c *gin.Context) {
	code := c.Param("code")
	questions := s.service.GetQuestionsList(code)

	if len(questions) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Вопросы для темы не найдены"})
		return
	}

	c.IndentedJSON(http.StatusOK, questions)
}
