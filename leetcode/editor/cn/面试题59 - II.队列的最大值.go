package cn

import "container/list"

//请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都
//是O(1)。
//
// 若队列为空，pop_front 和 max_value 需要返回 -1
//
// 示例 1：
//
// 输入:
//["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
//[[],[1],[2],[],[],[]]
//输出: [null,null,null,2,1,2]
//
//
// 示例 2：
//
// 输入:
//["MaxQueue","pop_front","max_value"]
//[[],[],[]]
//输出: [null,-1,-1]
//
//
//
//
// 限制：
//
//
// 1 <= push_back,pop_front,max_value的总操作数 <= 10000
// 1 <= value <= 10^5
//
// Related Topics 栈 Sliding Window

//leetcode submit region begin(Prohibit modification and deletion)
type MaxQueue struct {
	maxq *list.List
	q    *list.List
}

func ConstructorMaxQueue() MaxQueue {
	return MaxQueue{
		maxq: list.New(),
		q:    list.New(),
	}
}

func (this *MaxQueue) Max_value() int {
	element := this.maxq.Front()
	if element == nil {
		if this.q.Front() != nil {
			return this.q.Front().Value.(int)
		}
		return -1
	}
	return element.Value.(int)
}

func (this *MaxQueue) Push_back(value int) {
	this.q.PushBack(value)
	for this.maxq.Back() != nil && value > this.maxq.Back().Value.(int) {
		this.maxq.Remove(this.maxq.Back())
	}
	this.maxq.PushBack(value)

}

func (this *MaxQueue) Pop_front() int {
	element := this.q.Front()
	if element == nil {
		return -1
	}
	this.q.Remove(element)
	if this.maxq.Front().Value == element.Value {
		this.maxq.Remove(this.maxq.Front())
	}
	return element.Value.(int)
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
//leetcode submit region end(Prohibit modification and deletion)
