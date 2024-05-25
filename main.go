package main

import (
	"fmt"
)

func calculator() {
	var a, b float64
	var op string

	fmt.Println("Введите первое число:")
	fmt.Scanln(&a)
	fmt.Println("Введите оператор (+, -, *, /):")
	fmt.Scanln(&op)
	fmt.Println("Введите второе число:")
	fmt.Scanln(&b)

	switch op {
	case "+":
		fmt.Println("Результат:", a+b)
	case "-":
		fmt.Println("Результат:", a-b)
	case "*":
		fmt.Println("Результат:", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Результат:", a/b)
		} else {
			fmt.Println("Ошибка: деление на ноль")
		}
	default:
		fmt.Println("Ошибка: неверный оператор")
	}
}

func main() {
	calculator()
}
