package cn

import (
	"fmt"
	"testing"
)

//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。 
//
// 说明: 
//
// 
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。 
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。 
// 
//
// 示例: 
//
// 输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6] 
// Related Topics 数组 双指针

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j = 0, 0
	for i < m+n && j < n {
		// 随着合并增多, 需要动态增加 i的边界. m里的值还没有完全检索完成
		if i < m + j && nums2[j] >= nums1[i]{
			i++
			continue
		}
		// 向后移动需要的元素. 逐渐剔除0
		for ti := m + n - 1; ti > i; ti-- {
			nums1[ti] = nums1[ti-1]
		}
		nums1[i] = nums2[j]
		// 两个指针均后移. 指向新的值.
		j++
		i++
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeSortedArray(t *testing.T) {
	var num1 = []int{0}
	merge(num1, 0, []int{1}, 1)
	fmt.Println(num1)
	num1 = []int{1}
	merge(num1, 1, []int{}, 0)
	fmt.Println(num1)
	num1 = []int{1, 2, 3, 0, 0, 0}
	merge(num1, 3, []int{2, 5, 6}, 3)
	fmt.Println(num1)
	num1 = []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	merge(num1, 6, []int{1, 2, 2}, 3)
	fmt.Println(num1)
	num1 = []int{-1, 0, 0, 3, 3, 3, 0, 0, 0, 0, 0}
	merge(num1, 6, []int{-2, -1, 1, 2, 2}, 5)
	fmt.Println(num1)
}
