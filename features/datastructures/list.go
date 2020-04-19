package datastructure

import (
	"container/list"
	"fmt"
)

func List_Sample01() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)
	intList.PushBack("---Hello World---")
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println("Element of the list -> ", element.Value)
	}
}
