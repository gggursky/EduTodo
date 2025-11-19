package service

import "github.com/Konstantin299/EduTodo.git/internal/models"

var question1 = models.Question{
	Name: "В каком году началось производсво Шуманет БМ?",
	Code: "1",
	Answers: []models.AnswerVariant{
		{
			Name:    "1999 год",
			Code:    "1",
			IsRight: false,
		},
		{
			Name:    "2000 год",
			Code:    "2",
			IsRight: true,
		},
		{
			Name:    "2005 год",
			Code:    "3",
			IsRight: false,
		},
	},
}
var question2 = models.Question{
	Name: "Какие противопожарные свойства у Шуманет БМ?",
	Code: "2",
	Answers: []models.AnswerVariant{
		{
			Name:    "КМ1",
			Code:    "1",
			IsRight: false,
		},
		{
			Name:    "НГ",
			Code:    "2",
			IsRight: true,
		},
		{
			Name:    "Н1",
			Code:    "3",
			IsRight: false,
		},
	},
}

var question3 = models.Question{
	Name: "Чье сырье используется для производства Шуманет БМ?",
	Code: "3",
	Answers: []models.AnswerVariant{
		{
			Name:    "Rockwool",
			Code:    "1",
			IsRight: true,
		},
		{
			Name:    "URSA",
			Code:    "2",
			IsRight: false,
		},
		{
			Name:    "ISOVER",
			Code:    "3",
			IsRight: false,
		},
	},
}

var question4 = models.Question{
	Name: "Толщина пергородки на сдвоенном каркасе 100 мм?",
	Code: "4",
	Answers: []models.AnswerVariant{
		{
			Name:    "268 мм",
			Code:    "1",
			IsRight: false,
		},
		{
			Name:    "158 мм",
			Code:    "2",
			IsRight: true,
		},
		{
			Name:    "168 мм",
			Code:    "3",
			IsRight: false,
		},
	},
}

var question5 = models.Question{
	Name: "Максимальная высота облицовки с применением креплений Виброфлекс-КС?",
	Code: "5",
	Answers: []models.AnswerVariant{
		{
			Name:    "6 м",
			Code:    "1",
			IsRight: false,
		},
		{
			Name:    "5,5 м",
			Code:    "2",
			IsRight: false,
		},
		{
			Name:    "10 м",
			Code:    "3",
			IsRight: true,
		},
	},
}

var questions = []models.Question{question1, question2, question3, question4, question5}
