package cn

import (
	"fmt"
	"sort"
	"testing"
)

//给定两个数组，编写一个函数来计算它们的交集。 
//
// 示例 1: 
//
// 输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2]
// 
//
// 示例 2: 
//
// 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [9,4] 
//
// 说明: 
//
// 
// 输出结果中的每个元素一定是唯一的。 
// 我们可以不考虑输出结果的顺序。 
// 
// Related Topics 排序 哈希表 双指针 二分查找

//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)
	var m = map[int]struct{}{}
	sort.Ints(nums1) // 使用排序实现双头指针遍历的可行性
	sort.Ints(nums2)
	var p1, p2 = 0, 0
	for p1 < l1 && p2 < l2 {
		if nums1[p1] == nums2[p2] {
			var el = nums2[p2]
			m[el] = struct{}{}
			p1++
			p2++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}

	var ints []int
	for k := range m {
		ints = append(ints, k)
	}
	return ints
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIntersectionOfTwoArrays(t *testing.T) {
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
