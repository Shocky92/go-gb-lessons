package main

import "fmt"

func main() {
	var name string

	fmt.Printf("Enter your name: \n")

	fmt.Scan(&name)
	fmt.Printf("Your name is %s.", name)
}
