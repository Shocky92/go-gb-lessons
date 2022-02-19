package main

import "fmt"

func fibonacci(num int) int {
	if (num == 0) || (num == 1) {
		return num
	}

	return fibonacci(num-1) + fibonacci(num-2)
}

func fibonacciMap(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}

	return fn[n]
}

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Println(fibonacci(num))
	fmt.Println(fibonacciMap(num))
}
