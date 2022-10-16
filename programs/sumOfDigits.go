package main

import "fmt"

// sumOfDigits calculates the sum of the digits that makes up a number
func sumOfDigits(number int) int {
	total := 0
	for number > 0 {
		total += number % 10
		number /= 10
	}
	return total
}

func main() {
	myNumber := 168
	fmt.Println(sumOfDigits(myNumber))
}
