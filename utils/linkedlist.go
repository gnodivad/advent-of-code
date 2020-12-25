package utils

import (
	"container/list"
	"fmt"
)

func PrintListAsIntArray(list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(int))
		if e.Next() != nil {
			fmt.Print(",")
		}
	}
	fmt.Println("")
}
