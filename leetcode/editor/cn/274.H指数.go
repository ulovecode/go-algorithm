package leetcode

import "sort"

//给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。
//
// h 指数的定义: “h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）”
//
//
//
// 示例:
//
// 输入: citations = [3,0,6,1,5]
//输出: 3
//解释: 给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
//     由于研究者有 3 篇论文每篇至少被引用了 3 次，其余两篇论文每篇被引用不多于 3 次，所以她的 h 指数是 3。
//
//
//
// 说明: 如果 h 有多种可能的值，h 指数是其中最大的那个。
// Related Topics 排序 哈希表

//leetcode submit region begin(Prohibit modification and deletion)
func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i]-citations[j] > 0
	})
	if len(citations) == 0 {
		return 0
	}
	var i = 0
	for i < len(citations) && citations[i] > i {
		i++
	}
	return i
}

//func hIndex(citations []int) int {
//	n := len(citations)
//	bucket := make([]int, n+1)
//	for _, citation := range citations {
//		if citation >= n {
//			bucket[n] ++
//		} else {
//			bucket[citation] ++
//		}
//	}
//	var k = n
//	for s := bucket[n]; k > s; s += bucket[k] {
//		k--
//	}
//	return k
//}

//leetcode submit region end(Prohibit modification and deletion)
