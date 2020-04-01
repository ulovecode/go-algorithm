package main

import (
	"container/list"
)

//func main() {
//	var n = 0
//	fmt.Scanf("%d", &n)
//	stack := list.New()
//	for ; n != 0; n-- {
//		var v = 0
//		fmt.Scanf("%d", &v)
//		stack.PushFront(v)
//	}
//	reverse(stack)
//	for stack.Len() != 0 {
//		back := stack.Back()
//		fmt.Print(string(back.Value.(int))+" ")
//		stack.Remove(back)
//	}
//}

func reverse(stack *list.List) {
	if stack.Len() == 0 {
		return
	}
	element := stack.Back()
	stack.Remove(element)
	reverse(stack)
	putTopToDown(stack, element.Value.(int))
}

func putTopToDown(stack *list.List, value int) {
	if stack.Len() == 0 {
		stack.PushBack(value)
		return
	}
	element := stack.Back()
	stack.Remove(element)
	putTopToDown(stack, value)
	stack.PushBack(element.Value)
}
