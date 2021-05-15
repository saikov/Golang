package main

import (
	"fmt"
	"strings"
)

func main() {
	var chat string
	fmt.Print("Введите строку: ")
	fmt.Scan(&chat)
	if strings.Contains(chat, "чат") {
		fmt.Println("БОТ")
		return
	}
	fmt.Println("НЕ БОТ")
}
