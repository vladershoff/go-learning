/*
Создать новый срез, содержащий только элементы больше определённого значения.

Пример:
	[1, 5, 3, 8, 2, 9] с threshold=5 → [8, 9]
	nums := []int{1, 5, 3, 8, 2, 9, 4, 7, 12, 1, 4, 15, 16} threshold := 5
*/

package main

import (
	"fmt"
)

func newSlice(nums *[]int, threshold int) *[]int {
	// Используем указатели на слайсы, чтобы избежать копирования данных

	// Создаем новый слайс.
	// Чтобы при добавлении новых элементов не пересоздавался базовый массив
	// используем глубину равную глубине исходного слайса.
	res := make([]int, 0, len(*nums))

	// Перекладываем в новый слайс элементы больше порога
	for _, v := range *nums {
		if v > threshold {
			res = append(res, v)
		}
	}
	return &res
}

func main() {
	// Исходный слайс и порог
	nums := []int{1, 5, 3, 8, 2, 9, 4, 7, 12, 1, 4, 15, 16}
	threshold := 5
	newNums := newSlice(&nums, threshold)
	fmt.Println(*newNums)
}
