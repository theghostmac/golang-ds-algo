package main

import "fmt"

func main() {
	// arrays
	var myArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)
	fmt.Println(len(myArray))
	// slices
	var mySlice = []int{2, 4, 6, 8}
	fmt.Println(mySlice)
	// maps
	var languages = map[int]string{
		1: "English",
		2: "Japanese",
		3: "Igbo",
	}
	fmt.Println(languages)
}
