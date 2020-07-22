package leetcode

import "container/list"

//实现一个MyQueue类，该类用两个栈来实现一个队列。 示例： MyQueue queue = new MyQueue(); queue.push(1);
//queue.push(2); queue.peek();  // 返回 1 queue.pop();   // 返回 1 queue.empty(); // 返
//回 false 说明： 你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size 和 is empty
// 操作是合法的。 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。 假设所有操作都是有效的
//（例如，一个空的队列不会调用 pop 或者 peek 操作）。 Related Topics 栈

//leetcode submit region begin(Prohibit modification and deletion)
type MyQueue struct {
	s1 *list.List
	s2 *list.List
}

/** Initialize your data structure here. */
func ConstructorMyQueue() MyQueue {
	return MyQueue{
		s1: list.New(),
		s2: list.New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.fullValueS2()
	element := this.s2.Back()
	this.s2.Remove(element)
	return element.Value.(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.fullValueS2()
	return this.s2.Back().Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s1.Len() == 0 && this.s2.Len() == 0
}

func (q *MyQueue) fullValueS2() {
	if q.s2.Len() == 0 {
		for q.s1.Len() != 0 {
			element := q.s1.Back()
			q.s2.PushBack(element.Value)
			q.s1.Remove(element)
		}
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
