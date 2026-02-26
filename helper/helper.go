package helper

import (
	"fmt"
	"math"
)

func Calculate() string {

	fmt.Println("Введите выражение в формате: 1 + 2")
	fmt.Println("Для выхода введите: 0 0 0")

	for {

		var firstNumber, secondNumber float64
		var operation string
		fmt.Scan(&firstNumber, &operation, &secondNumber)

		if firstNumber == 0 && operation == "0" && secondNumber == 0 {
			break
		}

		switch operation {
		case "+":
			fmt.Println("Ответ:", firstNumber+secondNumber)
		case "-":
			fmt.Println("Ответ:", firstNumber-secondNumber)
		case "*":
			fmt.Println("Ответ:", firstNumber*secondNumber)
		case "/":
			if secondNumber == 0 {
				fmt.Println("Ошибка: Деление на ноль")
				return operation
			}
			fmt.Printf("Ответ: %.2f\n", firstNumber/secondNumber)
		case "%":
			fmt.Println("Ответ:", int(firstNumber)%int(secondNumber))
		case "^":
			fmt.Println("Ответ:", math.Pow(float64(firstNumber), float64(secondNumber)))
		default:
			fmt.Println("Неверная операция")
		}
	}
	return "exit"
}

func GuessNumber() string {
	var number int
	fmt.Println("Введите число от 1 до 100")
	fmt.Scan(&number)
	if number == 42 {
		fmt.Println("Вы угадали число")
	} else {
		fmt.Println("Вы не угадали число")
	}
	return "exit"
}
