package dataStructures

import "fmt"

type Array struct {
	len      int
	elements []interface{}
}

// Get method receives an index of the array and returns an interface type.
func (a *Array) Get(index int) (interface{}, error) {
	if index > a.len-1 {
		return nil, fmt.Errorf("Illegal index.")
	}
	return a.elements[index], nil
}

// Push method receives a value and adds the value to the array, then returns the new length of the array.
func (a *Array) Push(value interface{}) int {
	a.elements = append(a.elements, value)
	a.len++
	return a.len
}

// Pop method removes the last element in the array
func (a *Array) Pop() (interface{}, error) {
	if a.len == 0 {
		return nil, fmt.Errorf("There is nothing to pop. The length of the array is 0.")
	}
	lastItem := a.elements[a.len-1]
	a.elements = a.elements[:a.len-1]
	a.len--
	return lastItem, nil
}

// Delete method removes the element at a provided index in the array, then shifts the succeeding elements to fill the space
func (a *Array) Delete(index int) (interface{}, error) {
	deletedItem, err := a.Get(index)
	if err != nil {
		return nil, err
	}
	a.shiftItems(index)
	_, err = a.Pop()
	if err != nil {
		return nil, err
	}
	return deletedItem, err
}

// shiftItems method helps the Delete method by reshuffling the array to avoid empty spaces
func (a *Array) shiftItems(index int) {
	for i := index; i < a.len-1; i++ {
		a.elements[i] = a.elements[i+1]
	}
}

// Length method returns the length of the array after relevant operations are ended
func (a *Array) Length() int {
	return a.len
}

// check array_test.go file for testing.
