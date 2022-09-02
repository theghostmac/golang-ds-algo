package dataStructures

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	// access all elements
	var i int
	for i = 0; i < len(arr); i++ {
		fmt.Println("Printing array elements: ", arr[i])
	}
	var value int
	for i, value = range arr {
		fmt.Println(" range ", value)
	}

	var mySlice = []int{2, 4, 6, 8, 10}
	mySlice = append(mySlice, 12, 14)
	fmt.Println("Capacity = ", cap(mySlice))
	fmt.Println("Length = ", len(mySlice))
}

