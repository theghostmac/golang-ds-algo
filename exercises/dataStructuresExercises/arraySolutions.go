package dataStructuresExercises

/*
Case 1:
Write a method that will return the sum of all the elements of the integer list,
given list as an input argument.
*/

// SumArray collects an array and sums all the elements in it, then provides the total.
func SumArray(elements []int) int {
	size := len(elements)
	total := 0
	for index := 0; index < size; index++ {
		total = total + elements[index]
	}
	return total
}

/*
Case 2:
Write a method, which will search a list for some given value.
*/

// SequentialSearch searches for a value in a given list
func SequentialSearch(elements []int, value int) bool {
	size := len(elements)
	for index := 0; index < size; index++ {
		if value == elements[index] {
			return true
		}
	}
	return false
}

/*
Case 3:
Implement a Binary Search Tree
*/

func BinarySearch(data []int, value int) bool {
	size := len(data)
	var mid int
	low := 0
	high := size - 1
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}
