package main

// Lists

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(22)
	intList.PushBack(33)
	fmt.Println(intList)
}
