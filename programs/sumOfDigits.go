package main

import "fmt"

func findSumOfDigits(num int) int {
	total := 0
	for num > 0 {
		total += num % 10
		num /= 10
	}
	return total
}

func main() {
	myNum := 230
	a := findSumOfDigits(myNum)
	fmt.Println(a)
}
