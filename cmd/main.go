package main

import (
	"fmt"
	"math"
)

const (
	OperationAddition = "+"
	OperationSubtract = "-"
	OperationMultiply = "*"
	OperationDivide   = "/"
	OperationPower    = "^"
)

func main() {
	var firstOperand, secondOperand, result float64
	var operator string

	fmt.Println("Введите оператор (+ - * / ^), первый и второй операнд:")
	fmt.Scan(&operator, &firstOperand, &secondOperand)

	switch operator {
	case OperationAddition:
		result = firstOperand + secondOperand
	case OperationSubtract:
		result = firstOperand - secondOperand
	case OperationMultiply:
		result = firstOperand * secondOperand
	case OperationDivide:
		if secondOperand == 0 {
			fmt.Println("Деление на ноль невозможно")
			return
		}
		result = firstOperand / secondOperand
	case OperationPower:
		result = math.Pow(firstOperand, secondOperand)
	default:
		fmt.Println("Неверно выбран оператор")
		return
	}

	fmt.Println("Результат: ", result)
}
