package service

import "github.com/Konstantin299/EduTodo.git/internal/models"

var noise1 = models.NoiseSource{
	Name: "МОП",
	Code: "1",
}

var noise2 = models.NoiseSource{
	Name: "топот соседей сверху",
	Code: "2",
}

var noise3 = models.NoiseSource{
	Name: "детский центр",
	Code: "4",
}

var noise4 = models.NoiseSource{
	Name: "Шум самолетов",
	Code: "4",
}

var noise5 = models.NoiseSource{
	Name: "автострада",
	Code: "5",
}

var noise6 = models.NoiseSource{
	Name: "трамвайная линия",
	Code: "6",
}

var noise7 = models.NoiseSource{
	Name: "внешние блоки кондиционеров, караоке и детского центра",
	Code: "7",
}

var noise8 = models.NoiseSource{
	Name: "Лает и прыгает собака",
	Code: "8",
}

var noise9 = models.NoiseSource{
	Name: "Караоке",
	Code: "9",
}

var noise10 = models.NoiseSource{
	Name: "Внешний блок кондиционера",
	Code: "10",
}

var noise11 = models.NoiseSource{
	Name: "VRV блок кондиционера",
	Code: "11",
}

var noise12 = models.NoiseSource{
	Name: "Тоннель метро",
	Code: "12",
}
var noises = []models.NoiseSource{noise1, noise2, noise3, noise4, noise5, noise6, noise7, noise8, noise9, noise10, noise11, noise12}
