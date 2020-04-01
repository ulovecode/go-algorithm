package main

import (
	"container/list"
)

type Queue interface {
	Poll() int
	Add(v int)
	Peek() int
}

type queue struct {
	s1 *list.List
	s2 *list.List
}

func NewQueue() *queue {
	return &queue{
		s1: list.New(),
		s2: list.New(),
	}
}

func (q queue) Poll() int {
	q.fullValueS2()
	element := q.s2.Back()
	q.s2.Remove(element)
	return element.Value.(int)
}

func (q queue) Add(v int) {
	q.s1.PushBack(v)
}

func (q queue) Peek() int {
	q.fullValueS2()
	return q.s2.Back().Value.(int)
}

func (q queue) fullValueS2() {
	if q.s2.Len() == 0 {
		for q.s1.Len() != 0 {
			element := q.s1.Back()
			q.s2.PushBack(element.Value)
			q.s1.Remove(element)
		}
	}
}

//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Split(bufio.ScanWords)
//	queue := NewQueue()
//	for ; scanner.Scan() && n != 0; n-- {
//		if scanner.Text() == "add" {
//			scanner.Scan()
//			atoi, err := strconv.Atoi(scanner.Text())
//			if err != nil {
//				panic(err)
//			}
//			queue.Add(atoi)
//		} else if scanner.Text() == "peek" {
//			fmt.Println(queue.Peek())
//		} else if scanner.Text() == "poll" {
//			queue.Poll()
//		}
//	}
//}
