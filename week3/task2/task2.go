package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Product struct {
	Name  string
	Price float64
}

type Cart struct {
	Items   map[string]int
	Catalog map[string]Product
}

// Добавляет товар в корзину
func (cart *Cart) Add(productName string, qty int) error {
	if _, exist := cart.Catalog[productName]; !exist {
		return errors.New("Не найден товар в каталоге: " + productName)
	}

	if qty <= 0 {
		return errors.New("Количество товаров должно быть положительным значением: " + strconv.Itoa(qty))
	}

	// Инициализируем map корзины
	if cart.Items == nil {
		cart.Items = make(map[string]int)
	}

	// Добавляем товар в корзину
	cart.Items[productName] += qty

	return nil
}

// Вычисляет сумму товара в корзине с учетом скидки
func (cart *Cart) Total(discount func(float64) float64) float64 {
	// Общая стоимость товаров в корзине
	sum := 0.0

	for productName, qty := range cart.Items {
		// Прверяем наличие товара в каталоге
		product, exist := cart.Catalog[productName]
		if !exist {
			continue
		}

		// Извлекаем стоимость товара из каталога и умножаем
		// на количество в корзине
		sum += product.Price * float64(qty)
	}

	// Применяем к стоимости скидку
	return discount(sum)
}

func main() {
	cart := Cart{
		Catalog: map[string]Product{
			"молоко": {Name: "молоко", Price: 45.0},
			"яблоко": {Name: "яблоко", Price: 20.0},
			"кофе":   {Name: "кофе", Price: 150.0},
		},
	}

	// Добавляем продукты в корзину
	cart.Add("молоко", 2)
	cart.Add("яблоко", 10)

	// Вычисляем итоговую стоимость с учетом скидки 35%
	total := cart.Total(func(sum float64) float64 { return (1 - 0.35) * sum })

	fmt.Printf("Итоговая стоимость - %.2f\n", total)
}
