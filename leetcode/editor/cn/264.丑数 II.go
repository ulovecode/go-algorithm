package cn

import (
	"container/list"
	"math"
)

//编写一个程序，找出第 n 个丑数。
//
// 丑数就是只包含质因数 2, 3, 5 的正整数。
//
// 示例:
//
// 输入: n = 10
//输出: 12
//解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
//
// 说明:
//
//
// 1 是丑数。
// n 不超过1690。
//
// Related Topics 堆 数学 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	q2 := list.New()
	q3 := list.New()
	q5 := list.New()
	q2.PushBack(1)
	var minNumber int
	var q *list.List
	for n != 0 {
		minNumber, q = getMinNumber(q2, q3, q5)
		q.Remove(q.Front())
		if q == q2 {
			q2.PushBack(minNumber * 2)
			q3.PushBack(minNumber * 3)
			q5.PushBack(minNumber * 5)
		}

		if q == q3 {
			q3.PushBack(minNumber * 3)
			q5.PushBack(minNumber * 5)
		}

		if q == q5 {
			q5.PushBack(minNumber * 5)
		}
		n--
	}
	return minNumber
}

func getMinNumber(q2 *list.List, q3 *list.List, q5 *list.List) (int, *list.List) {
	result := math.MaxInt32
	var q *list.List
	if q2.Front() != nil && result > q2.Front().Value.(int) {
		result = q2.Front().Value.(int)
		q = q2
	}
	if q3.Front() != nil && result > q3.Front().Value.(int) {
		result = q3.Front().Value.(int)
		q = q3
	}
	if q5.Front() != nil && result > q5.Front().Value.(int) {
		result = q5.Front().Value.(int)
		q = q5
	}
	return result, q
}

//leetcode submit region end(Prohibit modification and deletion)
