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

// Check обрабатывает ответы пользователя и возвращает результат проверки.
// @Summary Проверка ответов пользователя
// @Description Принимает список выбранных пользователем ответов (questionCode + answerCode) и возвращает количество правильных ответов.
// @Tags тесты
// @Accept json
// @Produce json
// @Param request body []models.UserAnswer true "Ответы пользователя"
// @Success 200 {object} map[string]string "result" "Correct answers: X/Y"
// @Failure 400 {object} map[string]string "error" "Invalid request"
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
