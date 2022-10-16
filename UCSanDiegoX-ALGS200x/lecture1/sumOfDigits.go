package main

import "fmt"

// Constraints: 0 <= x, b <= 9

// sumTwoDigits collects two integers and sum them
func sumTwoDigits(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Enter two digits: ")
	// receive input from user
	var firstNum int
	fmt.Scanln(&firstNum)
	var secondNum int
	fmt.Scanln(&secondNum)
	// sum two inputs
	fmt.Println(sumTwoDigits(firstNum, secondNum))
}
