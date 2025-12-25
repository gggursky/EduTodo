package service

import (
	"os"

	"github.com/Konstantin299/EduTodo.git/internal/models"
)

const (
	ThemeFramelessName = "Бескаркасная звукоизоляция стен и потолков"
	ThemeFramelessCode = "code1"

	ThemeAcousticsBasicsName = "Основы строительной акустики и защиты от шума"
	ThemeAcousticsBasicsCode = "code2"

	ThemeSoundproofingBasicsName = "Основы технологии звукоизоляции"
	ThemeSoundproofingBasicsCode = "code3"
)

var (
	thema1 models.ThemaFullInfo
	thema2 models.ThemaFullInfo
	thema3 models.ThemaFullInfo
)

func init() {
	// Функция для чтения файлов
	load := func(path string) string {
		data, err := os.ReadFile(path)
		if err != nil {
			return "Ошибка: файл не найден (" + path + ")"
		}
		return string(data)
	}

	thema1 = models.ThemaFullInfo{
		Info: models.Thema{
			Name: ThemeFramelessName,
			Code: ThemeFramelessCode,
		},
		InfoBlock: load("internal/data/tema1.txt"),
		Questions: []models.Question{
			{
				Name: "Что является обязательным элементом при монтаже панелей ЗИПС на стену или потолок?",
				Code: "1",
				Answers: []models.AnswerVariant{
					{
						Name:    "Панель крепится только через виброузлы, входящие в комплект",
						Code:    "1",
						IsRight: true,
					},
					{
						Name:    "Панель должна крепиться к стене через обычные саморезы",
						Code:    "2",
						IsRight: false,
					},
					{
						Name:    "Панель обязательно клеится к стене по всей площади",
						Code:    "3",
						IsRight: false,
					},
				},
			},
			{
				Name: "Какая система ЗИПС обеспечивает наибольший индекс дополнительной звукоизоляции ΔRw?",
				Code: "2",
				Answers: []models.AnswerVariant{
					{
						Name:    "ЗИПС-Вектор",
						Code:    "1",
						IsRight: false,
					},
					{
						Name:    "ЗИПС-Синема",
						Code:    "2",
						IsRight: true,
					},
					{
						Name:    "ЗИПС-Слим",
						Code:    "3",
						IsRight: false,
					},
				},
			},
		},
	}

	thema2 = models.ThemaFullInfo{
		Info: models.Thema{
			Name: ThemeAcousticsBasicsName,
			Code: ThemeAcousticsBasicsCode,
		},
		InfoBlock: load("internal/data/tema2.txt"),
		Questions: []models.Question{
			{
				Name: "Что является основной причиной того, что многослойные конструкции обеспечивают более высокую звукоизоляцию по сравнению с однослойными?",
				Code: "3",
				Answers: []models.AnswerVariant{
					{
						Name:    "Увеличение толщины внешнего слоя стены",
						Code:    "1",
						IsRight: false,
					},
					{
						Name:    "Использование только тяжелых материалов",
						Code:    "2",
						IsRight: false,
					},
					{
						Name:    "Наличие воздушной прослойки, работающей как «пружина»",
						Code:    "3",
						IsRight: true,
					},
				},
			},
			{
				Name: "Что в наибольшей степени улучшает защиту от ударного шума в конструкции «плавающего пола»?",
				Code: "4",
				Answers: []models.AnswerVariant{
					{
						Name:    "Увеличение количества звукопоглощающих материалов в стяжке",
						Code:    "1",
						IsRight: false,
					},
					{
						Name:    "Уменьшение динамической жесткости упругого слоя",
						Code:    "2",
						IsRight: true,
					},
					{
						Name:    "Увеличение расстояния между стенами и полом",
						Code:    "3",
						IsRight: false,
					},
				},
			},
		},
	}

	thema3 = models.ThemaFullInfo{
		Info: models.Thema{
			Name: ThemeSoundproofingBasicsName,
			Code: ThemeSoundproofingBasicsCode,
		},
		InfoBlock: load("internal/data/tema3.txt"),
		Questions: []models.Question{
			{
				Name: "Что произойдёт, если при монтаже звукоизоляционной конструкции не использовать вибропрокладки?",
				Code: "5",
				Answers: []models.AnswerVariant{
					{
						Name:    "Звукоизоляция улучшится за счёт более жёсткой конструкции",
						Code:    "1",
						IsRight: false,
					},
					{
						Name:    "Примыкания станут более герметичными",
						Code:    "2",
						IsRight: false,
					},
					{
						Name:    "Шум начнёт лучше передаваться через жёсткие соединения и снизит эффективность системы",
						Code:    "3",
						IsRight: true,
					},
				},
			},
			{
				Name: "Какой материал нельзя использовать в качестве поглощающего слоя внутри каркасной звукоизоляционной системы?",
				Code: "6",
				Answers: []models.AnswerVariant{
					{
						Name:    "Пенопласт",
						Code:    "1",
						IsRight: true,
					},
					{
						Name:    "Шуманет-БМ",
						Code:    "2",
						IsRight: false,
					},
					{
						Name:    "Шуманет-ЭКО",
						Code:    "3",
						IsRight: false,
					},
				},
			},
		},
	}
}

var Themas = []models.ThemaFullInfo{thema1, thema2, thema3}

var course = []models.Thema{
	{Name: `Бескаркасная звукоизоляция стен и потолков`, Code: `code1`},
	{Name: `Основы строительной акустики и защиты от шума`, Code: `code2`},
	{Name: `Основы технологии звукоизоляции`, Code: `code3`},
}
