package array

import "fmt"

type Array struct {
	len      int
	elements []interface{}
}

func (a *Array) Get(index int) (interface{}, error) {
	if index > a.len-1 {
		return nil, fmt.Errorf("Illegal index.")
	}
	return a.elements[index], nil
}
