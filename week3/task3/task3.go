package main

import (
	"fmt"
	"slices"
)

type Network struct {
	Friends map[string][]string
}

// Добавляет двух пользователей в друзья друг к другу
func (net *Network) AddFriend(user1, user2 string) {
	// Инициализируем map друзей
	if net.Friends == nil {
		net.Friends = make(map[string][]string)
	}

	// Слайз друзей
	var frds []string

	// Находим друзей первого пользователя и добавляем второго
	// пользователя в друзья, если он не был добавлен ранее
	frds = net.Friends[user1]
	if !slices.Contains(frds, user2) {
		net.Friends[user1] = append(frds, user2)
	}

	// Находим друзей второго пользователя и добавляем первого
	// пользователя в друзья, если он не был добавлен ранее
	frds = net.Friends[user2]
	if !slices.Contains(frds, user1) {
		net.Friends[user2] = append(frds, user1)
	}
}

func SameFriends(net *Network, user1, user2 string) []string {
	// Инициализируем map друзей
	if net.Friends == nil {
		net.Friends = make(map[string][]string)
	}

	frds1 := net.Friends[user1]
	frds2 := net.Friends[user2]

	res := make([]string, 0, len(frds1))
	for _, v := range frds1 {
		if slices.Contains(frds2, v) {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	net := Network{}

	net.AddFriend("Петров", "Иванов")
	net.AddFriend("Петров", "Смирнов")
	net.AddFriend("Петров", "Соколов")
	net.AddFriend("Иванов", "Смирнов")
	net.AddFriend("Иванов", "Соколов")
	net.AddFriend("Смирнов", "Соколов")

	frds := SameFriends(&net, "Петров", "Иванов")
	fmt.Println(frds)
}
