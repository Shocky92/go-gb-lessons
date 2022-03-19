package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/manifoldco/promptui"
)

const (
	OperationAddition = "Сложение +"
	OperationSubtract = "Вычитание -"
	OperationMultiply = "Умножение *"
	OperationDivide   = "Деление /"
	OperationPower    = "Возведение в степень ^"
	OperationModulo   = "Остаток от деления %"
	OperationMax      = "Максимальное число"
	OperationMin      = "Минимальное число"
)

func getOperator() (string, error) {
	prompt := promptui.Select{
		Label: "Выберите оператор:",
		Items: []string{
			OperationAddition,
			OperationSubtract,
			OperationMultiply,
			OperationDivide,
			OperationPower,
			OperationModulo,
			OperationMax,
			OperationMin,
		},
	}

	_, operator, err := prompt.Run()

	if err != nil {
		return "", errors.New("Ошибка! Неверно выбран оператор.")
	}

	return operator, nil
}

func calculate(operator string, firstOperand float64, secondOperand float64) (float64, error) {
	var result float64

	switch operator {
	case OperationAddition:
		result = firstOperand + secondOperand
	case OperationSubtract:
		result = firstOperand - secondOperand
	case OperationMultiply:
		result = firstOperand * secondOperand
	case OperationDivide:
		if secondOperand == 0 {
			return result, errors.New("Деление на ноль невозможно")
		}
		result = firstOperand / secondOperand
	case OperationPower:
		result = math.Pow(firstOperand, secondOperand)
	case OperationModulo:
		result = math.Mod(firstOperand, secondOperand)
	case OperationMax:
		result = math.Max(firstOperand, secondOperand)
	case OperationMin:
		result = math.Min(firstOperand, secondOperand)
	default:
		return result, errors.New("Неверно выбран оператор")
	}

	return result, nil
}

func main() {
	var operator string
	var firstOperand, secondOperand, result float64

	operator, _ = getOperator()

	fmt.Println("Введите первый и второй операнд:")
	fmt.Scan(&firstOperand, &secondOperand)

	result, _ = calculate(operator, firstOperand, secondOperand)

	fmt.Println("Результат: ", result)
}
