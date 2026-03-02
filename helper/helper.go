package helper

import (
	"fmt"
	"math"
	"math/rand"
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
	n := rand.Intn(100)
	var number int
	fmt.Println("Введите число от 1 до 100")
	for {
		fmt.Scan(&number)
		if n == number {
			fmt.Println("Вы угадали число")
			break
		} else if n < number {
			fmt.Println("Число меньше")
		} else if n > number {
			fmt.Println("Число больше")
		}
	}
	return "exit"
}
