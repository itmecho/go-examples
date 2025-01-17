package main

import "fmt"

/*
 * variadic.go program will demonstrate use of
 * variadic paramters in Go programming language
 */
func main() {
	//passing zero paramter
	printNumbers()

	//passing two parameters
	printNumbers(5, 6)

	//passing slice
	slice := []int{1, 2}
	printNumbers(slice...)
}

/*
 * This function can take zero or more paramters
 */
func printNumbers(numbers ...int) {
	for index, number := range numbers {
		fmt.Printf("Number at index %d is %d", index, number)
		fmt.Println()
	}
}
