package main

import (
	"container/list"
)

type minStack struct {
	stack    *list.List
	minStack *list.List
}

func NewMinStack() *minStack {
	return &minStack{
		stack:    list.New(),
		minStack: list.New(),
	}
}

func (m *minStack) Pop() int {
	element := m.stack.Back()
	m.stack.Remove(element)
	if element.Value == m.minStack.Back().Value {
		m.minStack.Remove(m.minStack.Back())
	}
	return element.Value.(int)
}

func (m *minStack) Push(v int) {
	m.stack.PushBack(v)
	if m.minStack.Len() != 0 {
		if getBackValue(m.minStack) >= v {
			m.minStack.PushBack(v)
		}
	} else {
		m.minStack.PushBack(v)
	}
}

func (m *minStack) GetMin() int {
	return m.minStack.Back().Value.(int)
}

func getBackValue(s *list.List) int {
	return s.Back().Value.(int)
}

//
//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Split(bufio.ScanWords)
//	newMinStack := NewMinStack()
//	for ; scanner.Scan() && n != 0; n-- {
//		if scanner.Text() == "push" {
//			scanner.Scan()
//			atoi, err := strconv.Atoi(scanner.Text())
//			if err != nil {
//				panic(err)
//			}
//			newMinStack.Push(atoi)
//		} else if scanner.Text() == "getMin" {
//			fmt.Println(newMinStack.GetMin())
//		} else if scanner.Text() == "pop" {
//			newMinStack.Pop()
//		}
//	}
//}
