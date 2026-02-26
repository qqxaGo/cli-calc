package main

import (
	"fmt"

	"github.com/qqxaGo/cli-calc/helper"
)

func main() {
	var choose string
	fmt.Println("введите: 'Guess' или 'Calc'")
	fmt.Scan(&choose)

	switch choose {
	case "Guess":
		helper.GuessNumber()
	case "Calc":
		helper.Calculate()
	default:
		fmt.Println("Неверная команда")
	}
}
