//给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
//
// 示例 1:
//
// 输入: [1,2,0]
//输出: 3
//
//
// 示例 2:
//
// 输入: [3,4,-1,1]
//输出: 2
//
//
// 示例 3:
//
// 输入: [7,8,9,11,12]
//输出: 1
//
//
// 说明:
//
// 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
//

package cn

import "math"

func firstMissingPositive(nums []int) int {
	var contain1 bool
	for _, v := range nums {
		if v == 1 {
			contain1 = true
		}
	}
	if !contain1 {
		return 1
	}
	length := len(nums)
	if length == 1 {
		return 2
	}

	for i, v := range nums {
		if v > length || v <= 0 {
			nums[i] = 1
		}
	}

	for _, v := range nums {
		v = int(math.Abs(float64(v)))
		if v == length {
			nums[0] = -int(math.Abs(float64(nums[0])))
		} else {
			nums[v] = -int(math.Abs(float64(nums[v])))
		}
	}
	for i, v := range nums {
		if v > 0 {
			if i == 0 {
				continue
			}
			return i
		}
	}
	if nums[0] > 0 {
		return length
	}
	return length + 1
}
