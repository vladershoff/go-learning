package main

import (
	"fmt"
)

type Suspect struct {
	Name      string
	Height    int
	HairColor string
	HasScars  bool
}

func searchSuspect(suspects *[]Suspect) {
	// Используем указатель на слайс, чтобы избежать копирования данных

	const (
		targetHeight    = 185      // Фильтр по росту
		heightTolerance = 5        // Диапазон для фильтра по росту
		targetHair      = "blonde" // Фильтр по цвету волос
		targetScars     = true     // Фильтр по наличию шрама
	)

	var (
		minHeight = targetHeight - heightTolerance // Минимальный рост
		maxHeight = targetHeight + heightTolerance // Максимальный рост
	)

	for _, v := range *suspects {
		// Исключаем подозреваемых по росту с учетом диапазона
		if v.Height < minHeight && maxHeight < v.Height {
			continue
		}

		// Исключаем подозреваемых по цвету волос
		if v.HairColor != targetHair {
			continue
		}

		// Исключаем подозреваемых по наличию шраму
		if v.HasScars != targetScars {
			continue
		}

		// Выводим подозреваемого прошедшего через фильтры
		fmt.Printf("%s, %d, %s, %t\n", v.Name, v.Height, v.HairColor, v.HasScars)
	}
}

func main() {
	suspects := []Suspect{
		{
			Name:      "Nickolay",
			Height:    180,
			HairColor: "blonde",
			HasScars:  false,
		}, {
			Name:      "Andrew",
			Height:    175,
			HairColor: "brown",
			HasScars:  true,
		}, {
			Name:      "Alexey",
			Height:    185,
			HairColor: "blonde",
			HasScars:  true,
		}, {
			Name:      "Mikhail",
			Height:    170,
			HairColor: "white",
			HasScars:  false,
		}, {
			Name:      "Vitalii",
			Height:    182,
			HairColor: "blonde",
			HasScars:  true,
		},
	}

	searchSuspect(&suspects)
}
