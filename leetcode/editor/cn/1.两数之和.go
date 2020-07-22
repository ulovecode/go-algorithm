//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
//
package leetcode

// 在进行迭代并将元素插入到表中的同时，我们还会回过头来检查表中是否已经存在当前元素所对应的目标元素。如果它存在，那我们已经找到了对应解，并立即将其返回。
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if i2, ok := numsMap[v]; ok && i != i2 {
			return []int{i2, i}
		}
		numsMap[target-v] = i
	}
	return nil
}

// 复杂度分析：
//
//时间复杂度：O(n)O(n)，
//我们只遍历了包含有 nn 个元素的列表一次。在表中进行的每次查找只花费 O(1)O(1) 的时间。
//
//空间复杂度：O(n)O(n)，
//所需的额外空间取决于哈希表中存储的元素数量，该表最多需要存储 nn 个元素。
//
//作者：LeetCode
//链接：https://leetcode-cn.com/problems/two-sum/solution/liang-shu-zhi-he-by-leetcode-2/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
