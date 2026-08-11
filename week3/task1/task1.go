package main

import (
	"errors"
	"fmt"
)

type Hero struct {
	Name      string
	Power     int
	Inventory map[string]int
}

// Добавляет предмет в инвентарь героя
func (h *Hero) AddItem(item string, count int) error {
	// Инициализируем map инвентаря
	if h.Inventory == nil {
		h.Inventory = make(map[string]int)
	}

	if count <= 0 {
		return errors.New("Количество инвентаря должно быть положительное " + item)
	}

	// Добавляем инвентарь
	h.Inventory[item] += count

	return nil
}

// Использует(тратит) инвентарь героя
func (h *Hero) UseItem(item string, count int) error {
	// Инициализируем map инвентаря
	if h.Inventory == nil {
		h.Inventory = make(map[string]int)
	}

	if count <= 0 {
		return errors.New("Количество инвентаря должно быть положительное " + item)
	}

	if h.Inventory[item] < count {
		return errors.New("Недостаточное количество инвентаря " + item)
	}

	// Расходуем инвентарь
	h.Inventory[item] -= count

	// Удаляем инвентарь из map, если он закончился
	if h.Inventory[item] == 0 {
		delete(h.Inventory, item)
	}

	return nil
}

// Выводит текущие предметы инвенторя
func (h *Hero) Summary() {
	for k, v := range h.Inventory {
		fmt.Printf("%s - %d\n", k, v)
	}
}

func main() {
	h := Hero{Name: "User", Power: 100}

	// Добавляем предметы в инвентарь
	h.AddItem("Bow", 50)
	h.AddItem("Bow", 20)
	h.AddItem("Helmet", 30)
	h.AddItem("Shield", 50)

	var err error

	// Используем лук
	err = h.UseItem("Bow", 60)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// Используем шлем
	err = h.UseItem("Helmet", 40)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// Выводим оставшиеся предметы инвентаря
	h.Summary()
}
