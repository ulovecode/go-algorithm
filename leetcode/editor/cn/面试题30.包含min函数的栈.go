package leetcode

import "container/list"

//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
//
//
//
// 示例:
//
// MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.min();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.min();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// 各函数的调用总次数不超过 20000 次
//
//
//
//
// 注意：本题与主站 155 题相同：https://leetcode-cn.com/problems/min-stack/
// Related Topics 栈 设计

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	stack    *list.List
	minStack *list.List
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{
		stack:    list.New(),
		minStack: list.New(),
	}
}

func (m *MinStack) Push(v int) {
	m.stack.PushBack(v)
	if m.minStack.Len() != 0 {
		if getBackValue(m.minStack) >= v {
			m.minStack.PushBack(v)
		}
	} else {
		m.minStack.PushBack(v)
	}
}

func (m *MinStack) Pop() {
	element := m.stack.Back()
	m.stack.Remove(element)
	if element.Value == m.minStack.Back().Value {
		m.minStack.Remove(m.minStack.Back())
	}
}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}

func (m *MinStack) Min() int {
	return m.minStack.Back().Value.(int)
}

func getBackValue(s *list.List) int {
	return s.Back().Value.(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
//leetcode submit region end(Prohibit modification and deletion)
