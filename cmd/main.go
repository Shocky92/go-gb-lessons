package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getNumbers() []int64 {
	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)

	enter := "Введите число или нажмите Enter чтобы получить результат:"
	for fmt.Println(enter); scanner.Scan(); fmt.Println(enter) {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		inputNums = append(inputNums, num)
	}

	return inputNums
}

func sortNumbers(nums ...int64) []int64 {
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		j := i
		for ; j >= 1 && nums[j-1] > x; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = x
	}

	return nums
}

func main() {
	sorted := sortNumbers(getNumbers()...)

	for i := 1; i < len(sorted); i++ {
		fmt.Println(sorted[i])
	}
}
