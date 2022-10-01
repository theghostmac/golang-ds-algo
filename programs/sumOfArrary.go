package main

import "fmt"

// findSumOfArray calculates the sum of elements in an array
func findSumOfArray(arr []int) int {
	total := 0
	size := len(arr)
	for i := 0; i < size; i++ {
		total += arr[i]
	}
	return total
}

func main() {
	myArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	show := findSumOfArray(myArr)
	fmt.Println(show)
}
